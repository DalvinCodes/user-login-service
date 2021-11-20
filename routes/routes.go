package routes

import (
	"user-login-service/controller"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes (app * fiber.App, ctrl controller.UserController) {
	authRoutes := app.Group("/auth")
	publicRoutes := app.Group("/public")
	AuthRoutes(authRoutes, ctrl)
	PublicRoutes(publicRoutes, ctrl)
}
