package routes

import (
	"user-login-service/controller"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(r fiber.Router, ctrl controller.UserController) {
	UserPublicRoutes(r, ctrl)
	//TODO: Implement Public Routes
}
