package controller

import (
	helper "user-login-service/controller/utilities"
	"user-login-service/domain"
	"user-login-service/domain/dto"
	"user-login-service/service"

	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	Create(ctx *fiber.Ctx) error
	GetByID(tx *fiber.Ctx) error
	GetByUsername(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type controller struct {
	us service.UserService
}

func NewUserController(us service.UserService) UserController {
	return &controller{us: us}
}

func (c *controller) Create(ctx *fiber.Ctx) error {
	// initializes userDTO
	var userDTO dto.UserDTO 

	// takes JSON payload, binds it to model, and validate for error
	if err := helper.ParseAndValidatePayload(ctx, &userDTO); err != nil {
		return err
	}

	// maps DTOs to models
	userResponse, user := dto.UserDTOMapper(&userDTO)

	// creates a new user
	if err := c.us.Create(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	// returns the created userResponseDTO
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": userResponse,
	})
}

func (c *controller) GetByID(ctx *fiber.Ctx) error {
	//takes id field from parameters
	id := ctx.Query("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty or invalid id",
		})
	}

	//searches for ID in userDB
	user, err := c.us.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "unable to retrieve user",
			"error":   err,
		})
	}
	
	//maps user from db to DTO
	dto := dto.MapUserModelToDTO(user)

	//returns user
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dto,
	})

}

func (c *controller) GetByUsername(ctx *fiber.Ctx) error {
	
	username := ctx.Query("username")
	
	if username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty or invalid username",
		})
	}

	user, err := c.us.GetByUsername(username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user",
		})
	}

	userResponse := dto.MapUserModelToDTO(user)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": userResponse,
	})
}

func (c *controller) Delete(ctx *fiber.Ctx) error {
	id := ctx.Query("id")

	user, err := c.us.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error deleting user",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusNoContent).JSON(&user)

}

func (c *controller) Update(ctx *fiber.Ctx) error {
	var user *domain.User

	id := ctx.Query("id")

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	user.ID = id
	if err := c.us.Update(id, user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user updated",
	})

}
