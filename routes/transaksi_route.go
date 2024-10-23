package routes

import (
	"github.com/adieos/ets-pweb-be/controller"
	"github.com/adieos/ets-pweb-be/middleware"
	"github.com/adieos/ets-pweb-be/service"
	"github.com/gin-gonic/gin"
)

func Transaksi(route *gin.Engine, transaksiController controller.TransaksiController, jwtService service.JWTService) {
	routes := route.Group("/api/transaksi")
	{
		// Transaksi
		routes.POST("", middleware.Authenticate(jwtService), transaksiController.Create)
		routes.GET("", middleware.Authenticate(jwtService), transaksiController.GetAll)
		routes.GET("/komunal", middleware.Authenticate(jwtService), transaksiController.GetAllKomunal)
		routes.GET("/:id", middleware.Authenticate(jwtService), transaksiController.GetById)
		routes.PUT("/:id", middleware.Authenticate(jwtService), transaksiController.Update)
		routes.DELETE("/:id", middleware.Authenticate(jwtService), transaksiController.Delete)
	}
}
