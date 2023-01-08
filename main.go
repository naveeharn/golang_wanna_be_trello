package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/config"
	"github.com/naveeharn/golang_wanna_be_trello/controller"
	"github.com/naveeharn/golang_wanna_be_trello/middleware"
	"github.com/naveeharn/golang_wanna_be_trello/repository"
	"github.com/naveeharn/golang_wanna_be_trello/service"
	"gorm.io/gorm"
)

var (
	// gorm database
	db *gorm.DB = config.SetupDatabaseConnection()

	// repositories
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	teamRepository repository.TeamRepository = repository.NewTeamRepository(db)

	// services
	authService service.AuthService = service.NewAuthService(userRepository)
	userService service.UserService = service.NewUserService(userRepository)
	jwtService  service.JWTService  = service.NewJWTService()
	teamService service.TeamService = service.NewTeamService(teamRepository)

	// controllers
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService)
	teamController controller.TeamController = controller.NewTeamController(teamService)
)

func main() {
	defer config.CloseDatabaseConnection(db)

	routers := gin.Default()
	routers.Use(cors.Default())

	routers.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	authRoutes := routers.Group("api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
	}

	userRoutes := routers.Group("api/user")
	{
		userRoutes.GET("/", middleware.AuthorizeJWT(jwtService, userService), userController.Profile)
		userRoutes.GET("/:id", middleware.AuthorizeJWT(jwtService, userService), userController.GetUserById)
		userRoutes.PUT("/", middleware.AuthorizeJWT(jwtService, userService), userController.UpdateUser)
		userRoutes.PUT("/reset-password", middleware.AuthorizeJWT(jwtService, userService), userController.ResetPassword)

	}

	teamRoutes := routers.Group("api/team")
	{
		teamRoutes.POST("/", middleware.AuthorizeJWT(jwtService, userService), teamController.CreateTeam)
	}

	routers.Run(":4011")

}
