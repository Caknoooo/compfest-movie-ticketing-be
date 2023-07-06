package routes

import (
	"github.com/Caknoooo/golang-clean_template/controller"
	"github.com/Caknoooo/golang-clean_template/middleware"
	"github.com/Caknoooo/golang-clean_template/services"
	"github.com/gin-gonic/gin"
)

func Router(route *gin.Engine, UserController controller.UserController, SeederController controller.SeederController, topupController controller.TopupController, jwtService services.JWTService) {
	routes := route.Group("/api/user")
	{
		routes.POST("", UserController.RegisterUser)
		routes.GET("", middleware.Authenticate(jwtService), UserController.GetAllUser)
		routes.POST("/login", UserController.LoginUser)
		routes.DELETE("/", middleware.Authenticate(jwtService), UserController.DeleteUser)
		routes.PUT("/", middleware.Authenticate(jwtService), UserController.UpdateUser)
		routes.GET("/me", middleware.Authenticate(jwtService), UserController.MeUser)
	}

	seederRoutes := route.Group("/api/seeder")
	{
		seederRoutes.GET("/", SeederController.GetAllBank)
		seederRoutes.GET("/:id", SeederController.GetBankByID)
	}

	topupRoutes := route.Group("/api/topup") 
	{
		topupRoutes.POST("/", middleware.Authenticate(jwtService), topupController.CreateTopup)
		topupRoutes.GET("/", middleware.Authenticate(jwtService), topupController.GetAllTopupUser)
		topupRoutes.GET("/:id", topupController.GetTopupByID)
	}
}