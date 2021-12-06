package helpers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

//
const (
	PrivateKeyFileName = "private-key.pem"
	PublicKeyFileName  = "public-key.pem"
)

func ErrBadRequestResponse(err error) error {
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}
	return nil
}

func ErrInternalServerResponse(err error) error {
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}
	return nil
}

func ErrUnauthorizedResponse(err error) error {
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrUnauthorized.Code,
			Message: err.Error(),
		}
	}
	return nil
}

func TokenErrorValidationResponseHandler(vErrs error) error {


	switch vErrs.(type) {
	case nil:
		return nil
	case *jwt.ValidationError:
		switch vErrs.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: "expired token",
			}
		case jwt.ValidationErrorNotValidYet:
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: "token is not valid",
			}
		case jwt.ValidationErrorMalformed:
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: "token is not authorized",
			}
		case jwt.ValidationErrorUnverifiable:
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: "token is not valid",
			}
		case jwt.ValidationErrorSignatureInvalid:
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: "token is not valid",
			}
		default:
			return &fiber.Error{
				Code:    fiber.ErrUnauthorized.Code,
				Message: "token is not valid",
			}
		}
	}
	return nil
}

func ValidateBearerToken(bearerToken string) (string, error) {

	if bearerToken == "" {
		return "", &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "token is required",
		}
	}

	bearerToken = strings.TrimSpace(bearerToken)
	tempToken := strings.Split(bearerToken, " ")

	var token string

	if len(tempToken) == 2 {
		if tempToken[0] != "Bearer" {
			return "", &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: "invalid token type",
			}
		}
		token = tempToken[1]
	} else {
		return "", &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "invalid token",
		}
	}

	return token, nil
}
