package routes

import (
	"github.com/adieos/ets-pweb-be/controller"
	"github.com/adieos/ets-pweb-be/middleware"
	"github.com/adieos/ets-pweb-be/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userController controller.UserController, jwtService service.JWTService) {
	routes := route.Group("/api/user")
	{
		// User
		routes.POST("", userController.Register)
		routes.POST("/login", userController.Login)
		routes.GET("/me", middleware.Authenticate(jwtService), userController.Me)
	}
}
