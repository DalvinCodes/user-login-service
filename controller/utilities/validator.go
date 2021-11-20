package controller

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

//ParseAndValidatePayload binds json data to a struct 
//and validates payload against validation constraints
func ParseAndValidatePayload(ctx *fiber.Ctx, payload interface{}) *fiber.Error {
	if err := parsePayload(ctx, payload); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return validatePayload(payload)
}

//parsePayload binds json data to a struct
func parsePayload(ctx *fiber.Ctx, payload interface{}) *fiber.Error {
	if err := ctx.BodyParser(payload); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}

//validatePayload validates json data input against validation constraints
func validatePayload(payload interface{}) *fiber.Error {
	var errSlice []string
	var errString string
	err := Validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errSlice = append(errSlice, 
				fmt.Sprintf("%v doesn't satisfy the \"%v\" constraint\n", strings.ToLower(err.Field()), err.Tag()),
			)
		}
			errString = strings.Join(errSlice, ",")
			_errors := strings.Split(errString, ",")
			errors := fmt.Sprintln(strings.Trim(fmt.Sprint(_errors), "[]"))
			
			return &fiber.Error{
				Code:    fiber.ErrBadRequest.Code,
				Message: errors,
			}
		}
	return nil
}