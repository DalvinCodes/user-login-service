package main

import (
	"log"
	"user-login-service/controller"
	"user-login-service/db"
	"user-login-service/pkg"
	"user-login-service/repository"
	"user-login-service/routes"
	"user-login-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	// set up database
	DB, err := db.SetupPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	pwdService := pkg.NewPasswordService()
	repo := repository.NewUserRepository(DB)
	srv := service.NewUserService(repo, pwdService)
	ctrl := controller.NewUserController(srv)
	routes.LoadRoutes(app, ctrl)

	log.Fatal(app.Listen(":8080"))
}
