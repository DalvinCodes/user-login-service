package controller

import (
	"io/ioutil"
	"time"
	helper "user-login-service/controller/utilities"
	"user-login-service/domain/dto"
	"user-login-service/pkg/helpers"
	"user-login-service/pkg/middleware"
	"user-login-service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserController interface {
	Create(ctx *fiber.Ctx) error
	GetByID(tx *fiber.Ctx) error
	GetByUsername(ctx *fiber.Ctx) error
	UpdatePassword(ctx *fiber.Ctx) error
	AddAddress(ctx *fiber.Ctx) error
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
	user := dto.UserDTOMapper(&userDTO)

	// creates a new user
	responseUserJWT, err := c.us.Create(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	ctx.Set("id", responseUserJWT.ID)

	if err := middleware.GenerateToken(ctx); err != nil {
		return err
	}

	//removes id from header response
	ctx.Response().Header.Del("id")

	// returns the created userResponseDTO
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created",
	})
}

func (c *controller) GetByID(ctx *fiber.Ctx) error {
	//takes id field from query parameters
	id := ctx.Query("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty or invalid id",
		})
	}

	//searches for ID in userDB and returns an user or an error
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
	//takes username field from query parameters
	username := ctx.Query("username")

	if username == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "empty or invalid username",
		})
	}

	//searches for username in db and returns a user of error
	user, err := c.us.GetByUsername(username)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user",
		})
	}

	//maps user object to a DTO
	userResponse := dto.MapUserModelToDTO(user)

	// returns userDTO
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": userResponse,
	})
}

func (c *controller) UpdatePassword(ctx *fiber.Ctx) error {
	//initializes UserUpdatePasswordDTO
	var userReg *dto.UserUpdatePasswordDTO

	//takes id from query parameters
	id := ctx.Query("id")

	//takes JSON payload and parses it to userReg
	if err := ctx.BodyParser(&userReg); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	//updates the password with the associated ID
	if err := c.us.UpdatePassword(id, userReg); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	//returns a 200, status success if password updated in db
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user password updated",
	})

}

func (c *controller) AddAddress(ctx *fiber.Ctx) error {
	var addressDTO []dto.AddressDTO

	if err := ctx.BodyParser(&addressDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id := ctx.Query("id")

	if err := c.us.AddAddress(id, addressDTO); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user address updated",
	})
}

func (c *controller) Delete(ctx *fiber.Ctx) error {
	//takes username id form the query parameters
	id := ctx.Query("id")

	//returns an error if user is not found or deleted
	_, err := c.us.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error deleting user",
			"error":   err,
		})
	}

	//returns user deletion confirmation
	return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "user deleted",
	})

}

func GenerateToken(ctx *fiber.Ctx) error {

	//retrieves id string from request header and returns and error if empty
	id := ctx.Get("id")
	if id == "" {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "missing id string",
		}
	}

	//retrieves token expiration time to live from env file
	// var JWTSecretKeyExpirationTime = os.Getenv("JWT_SECRET_KEY_EXP")

	// //converts time string into an int
	// expiryTime, err := strconv.Atoi(JWTSecretKeyExpirationTime)
	// if err != nil {
	// 	return &fiber.Error{
	// 		Code:    fiber.ErrInternalServerError.Code,
	// 		Message: err.Error(),
	// 	}
	// }

	//reads in jwt private key from pem file and returns an error
	secretKeyFileString, err := ioutil.ReadFile(helpers.PrivateKeyFileName)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	//parse private key from pem file
	key, err := jwt.ParseRSAPrivateKeyFromPEM(secretKeyFileString)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	//generates an unsigned token
	token := jwt.New(jwt.SigningMethodRS256)

	token.Claims = &jwt.StandardClaims{
		Id:        id,
		Subject:   "User Creation and Authentication Service",
		Issuer:    "DalvinCodes",
		Audience:  "github.com/DalvinCodes/user-login-service",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(-1)).Unix(),
	}

	signedToken, err := token.SignedString(key)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": "Bearer " + signedToken,
	})
}
