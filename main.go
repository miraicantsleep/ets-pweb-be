package main

import (
	"log"
	"os"

	"github.com/adieos/ets-pweb-be/cmd"
	"github.com/adieos/ets-pweb-be/config"
	"github.com/adieos/ets-pweb-be/controller"
	"github.com/adieos/ets-pweb-be/middleware"
	"github.com/adieos/ets-pweb-be/repository"
	"github.com/adieos/ets-pweb-be/routes"
	"github.com/adieos/ets-pweb-be/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository      repository.UserRepository      = repository.NewUserRepository(db)
		transaksiRepository repository.TransaksiRepository = repository.NewTransaksiRepository(db)

		// Service
		userService      service.UserService      = service.NewUserService(userRepository, jwtService)
		transaksiService service.TransaksiService = service.NewTransaksiService(transaksiRepository, jwtService)

		// Controller
		userController      controller.UserController      = controller.NewUserController(userService)
		transaksiController controller.TransaksiController = controller.NewTransaksiController(transaksiService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)
	routes.Transaksi(server, transaksiController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
