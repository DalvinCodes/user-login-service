package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-login-service/controller"
)

func LoadRoutes (app * fiber.App, ctrl controller.UserController) {

	authRoutes := app.Group("/auth")
	AuthRoutes(authRoutes, ctrl)


	publicRoutes := app.Group("/public")
	PublicRoutes(publicRoutes, ctrl)
}
