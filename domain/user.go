package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          string     `json:"id"`
	UserType    uint       `json:"user_type" validate:"required,lt=5"`
	Username    string     `json:"username" validate:"required,min=5"`
	FirstName   string     `json:"first_name" validate:"required,min=3"`
	LastName    string     `json:"last_name" validate:"required,min=3"`
	Email       string     `json:"email" validate:"required,unique,email"`
	Password    string     `json:"password" validate:"required,min=6"`
	PhoneNumber string     `json:"phone_number" validate:"required,min=7"`
	Address     []Address `json:"address" validate:"required,dive,required"`
	ImageUrl    string     `json:"image_url,omitempty"`
	Active      bool       `json:"active"`
}
