package main

import (
	"log"
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
	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	teamRepository      repository.TeamRepository      = repository.NewTeamRepository(db)
	dashboardRepository repository.DashboardRepository = repository.NewDashboardRepository(db)
	noteRepository      repository.NoteRepository      = repository.NewNoteRepository(db)

	// services
	authService      service.AuthService      = service.NewAuthService(userRepository)
	userService      service.UserService      = service.NewUserService(userRepository)
	jwtService       service.JWTService       = service.NewJWTService()
	teamService      service.TeamService      = service.NewTeamService(teamRepository)
	dashboardService service.DashboardService = service.NewDashboardService(dashboardRepository)
	noteService      service.NoteService      = service.NewNoteService(noteRepository)

	// controllers
	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	userController      controller.UserController      = controller.NewUserController(userService)
	teamController      controller.TeamController      = controller.NewTeamController(teamService)
	dashboardController controller.DashboardController = controller.NewDashboardController(dashboardService)
	noteController      controller.NoteController      = controller.NewNoteController(noteService)
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

	userRoutes := routers.Group("api/user", middleware.AuthorizeJWT(jwtService, userService))
	{
		userRoutes.GET("/", userController.Profile)
		userRoutes.GET("/:userId", userController.GetUserById)
		userRoutes.PUT("/", userController.UpdateUser)
		userRoutes.PUT("/reset-password", userController.ResetPassword)

	}

	teamRoutes := routers.Group("api/team", middleware.AuthorizeJWT(jwtService, userService))
	{
		teamRoutes.POST("/", teamController.CreateTeam)
		teamRoutes.GET("/:teamId", teamController.GetTeamById)
		teamRoutes.GET("/", teamController.GetTeamsByOwnerUserId)
		teamRoutes.POST("/:teamId", teamController.AddMember)

		// teamRoutes.Group(":teamId/dashboard")
		// {
		// 	teamRoutes.POST("/", dashboardController.CreateDashboard)
		// 	teamRoutes.PUT("/:dashboardId", dashboardController.UpdateDashboard)
		// }
	}

	dashboardRoutes := routers.Group("api/team/:teamId/dashboard", middleware.AuthorizeJWT(jwtService, userService))
	{
		dashboardRoutes.POST("/", dashboardController.CreateDashboard)
		dashboardRoutes.PUT("/:dashboardId", dashboardController.UpdateDashboard)
	}

	noteRoutes := routers.Group("api/team/:teamId/dashboard/:dashboardId/note", middleware.AuthorizeJWT(jwtService, userService))
	{
		noteRoutes.POST("/", noteController.CreateNote)
		noteRoutes.PUT("/:noteId", noteController.UpdateNote)
		noteRoutes.DELETE("/:noteId", noteController.DeleteNote)
	}

	err := routers.Run(":4011")
	if err != nil {
		log.Fatal("Something went wrong")
	}

}
