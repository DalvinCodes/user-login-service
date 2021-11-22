package routes

import (
	"user-login-service/controller"

	"github.com/gofiber/fiber/v2"
)

func UserAdminRouter(ctrl controller.UserController, r fiber.Router) {
	v1 := r.Group("/v1")
	v1.Post("/user/create/", ctrl.Create)
	v1.Post("/user/updatepassword/", ctrl.UpdatePassword)
	v1.Get("/user/get/id/", ctrl.GetByID)
	v1.Get("/user/get/username/", ctrl.GetByUsername)
	v1.Delete("/user/delete/", ctrl.Delete)
}
