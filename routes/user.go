package routes

import (
	"github.com/gofiber/fiber/v2"
	"user-login-service/controller"
)

func UserAdminRouter(ctrl controller.UserController, r fiber.Router) {
	user := r.Group("/user")

	v1 := user.Group("/v1")
	v1.Post("/update/password/", ctrl.UpdatePassword)
	v1.Post("/update/address/add/", ctrl.AddAddress)
	v1.Get("/get/id/", ctrl.GetByID)
	v1.Get("/get/username/", ctrl.GetByUsername)
	v1.Delete("/delete/", ctrl.Delete)
}

func UserPublicRoutes(r fiber.Router, ctrl controller.UserController) {
	user := r.Group("/user")

	v1 := user.Group("/v1")
	v1.Post("/create/", ctrl.Create)
	v1.Post("/get/token/", controller.GenerateToken)

}
