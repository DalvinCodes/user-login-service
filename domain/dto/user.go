package dto

import (
	"time"
	"user-login-service/domain"
)

type UserDTO struct {
	UserType    uint          `json:"user_type" validate:"required,lt=5"`
	Username    string        `json:"username" validate:"required,min=5"`
	FirstName   string        `json:"first_name" validate:"required,min=5"`
	LastName    string        `json:"last_name" validate:"required,min=5"`
	Email       string        `json:"email" validate:"required,email"`
	Password    string        `json:"password" validate:"required,min=6"`
	PhoneNumber string        `json:"phone_number" validate:"required,min=7"`
	Address     []*AddressDTO `json:"address" validate:"required,dive,required"`
	ImageUrl    string        `json:"image_url,omitempty"`
}

type UserResponseDTO struct {
	ID          string        `json:"id" validate:"required"`
	UserType    uint          `json:"user_type" validate:"required,lt=5"`
	Username    string        `json:"username" validate:"required,min=5"`
	FirstName   string        `json:"first_name" validate:"required,min=5"`
	LastName    string        `json:"last_name" validate:"required,min=5"`
	Email       string        `json:"email" validate:"required,email"`
	PhoneNumber string        `json:"phone_number" validate:"required,min=7"`
	Address     []*AddressDTO `json:"address" validate:"required,dive,required"`
	ImageUrl    string        `json:"image_url,omitempty"`
	Active bool `json:"active" validate:"required"`
}

type UserCreatedDTO struct {
	CreateAt  time.Time `json:"created_at"`
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
}

func UserDTOMapper(dto *UserDTO) (*UserCreatedDTO, *domain.User) {

	var uAddresses []domain.Address
	for i := range dto.Address {
		var address domain.Address

		address.Line1 = dto.Address[i].Line1
		address.Line2 = dto.Address[i].Line2
		address.Line3 = dto.Address[i].Line3
		address.City = dto.Address[i].City
		address.State = dto.Address[i].State
		address.ZipCode = dto.Address[i].ZipCode
		address.Country = dto.Address[i].Country
		uAddresses = append(uAddresses, address)
	}

	user := &domain.User{
		UserType:    dto.UserType,
		Username:    dto.Username,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		Email:       dto.Email,
		Password:    dto.Password,
		PhoneNumber: dto.PhoneNumber,
		Address:     uAddresses,
		Active:      true,
	}

	responseDTO := &UserCreatedDTO{
		CreateAt:  time.Now(),
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}

	return responseDTO, user

}

func MapUserModelToDTO(user *domain.User) (userDTO UserResponseDTO){
	var dtoAddresses []*AddressDTO
	
	for i := range user.Address{
		var address AddressDTO

		address.Line1 = user.Address[i].Line1
		address.Line2 = user.Address[i].Line2
		address.Line3 = user.Address[i].Line3
		address.City = user.Address[i].City
		address.State = user.Address[i].State
		address.ZipCode = user.Address[i].ZipCode
		address.Country = user.Address[i].Country

		dtoAddresses = append(dtoAddresses, &address)
	}

	userDTO.ID = user.ID
	userDTO.UserType = user.UserType
	userDTO.Username = user.Username
	userDTO.FirstName = user.FirstName
	userDTO.LastName = user.LastName
	userDTO.Email = user.Email
	userDTO.PhoneNumber = user.PhoneNumber
	userDTO.Address = dtoAddresses
	userDTO.ImageUrl = user.ImageUrl
	userDTO.Active = user.Active


	return userDTO 
}
