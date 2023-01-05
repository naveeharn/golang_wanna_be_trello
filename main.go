package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/naveeharn/golang_wanna_be_trello/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabaseConnection(db)

	routers := gin.Default()
	routers.Use(cors.Default())

	routers.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello"})
	})

	routers.Run(":4011")

}
