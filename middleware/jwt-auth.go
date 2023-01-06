package middleware

import (
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/helper"
	"github.com/naveeharn/golang_wanna_be_trello/service"
)

func AuthorizeJWT(jwtService service.JWTService, userService service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || !strings.Contains(authHeader, "Bearer ") {
			response := helper.CreateErrorResponse("Failed to process request", "Token Authorization not found or not correct header pattern", helper.EmptyObj{})
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(strings.TrimPrefix(authHeader, "Bearer "))
		if err != nil {
			helper.LoggerErrorPath(runtime.Caller(0))
			log.Println(err.Error())
			response := helper.CreateErrorResponse("Token is not valid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			helper.LoggerErrorPath(runtime.Caller(0))
			log.Println(err.Error())
			response := helper.CreateErrorResponse("Token can not parse", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userId, ok := claims["userId"].(string)
		if !ok {
			response := helper.CreateErrorResponse("userId can not parse to string", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		// _, err = userService.GetUserById(userId)
		// if err != nil {
		// 	response := helper.CreateErrorResponse("userId from accessToken does not exists", err.Error(), nil)
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }
		ctx.Set("userId", userId)
	}
}
