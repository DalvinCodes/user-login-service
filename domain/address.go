package domain

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	ID      uint `json:"id,omitempty"`
	UserID  string
	Line1   string `json:"line_1" validate:"required"`
	Line2   string `json:"line_2,omitempty"`
	Line3   string `json:"line_3,omitempty"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
	Country string `json:"country" validate:"required"`
}
