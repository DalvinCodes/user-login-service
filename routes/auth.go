package routes

import (
	"user-login-service/controller"
	"user-login-service/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(authRouter fiber.Router, ctrl controller.UserController) {
	//TODO: Implement Auth Routes
	UserAdminRouter(ctrl, authRouter)
	authRouter.Use(middleware.ValidateToken)
}
