package routes

import (
	"user-login-service/controller"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(authRouter fiber.Router, ctrl controller.UserController) {
	//TODO: Implement Auth Routes
	UserAdminRouter(ctrl, authRouter)
}
