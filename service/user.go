package service

import (
	"errors"
	controller "user-login-service/controller/utilities"
	"user-login-service/domain"
	"user-login-service/domain/dto"
	"user-login-service/pkg"
	"user-login-service/repository"

	"github.com/google/uuid"
)

type UserService interface {
	Create(user *domain.User) (*dto.ResponseUserForJWT, error)
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	UpdatePassword(id string, userPasswordDTO *dto.UserUpdatePasswordDTO) error
	AddAddress(id string, address []dto.AddressDTO) error
	Delete(id string) (*domain.User, error)
}

type service struct {
	userRepo     repository.UserRepository
	passwordRepo pkg.PasswordRepo
}

func NewUserService(userRepo repository.UserRepository, passwordRepo pkg.PasswordRepo) UserService {
	return &service{
		userRepo:     userRepo,
		passwordRepo: passwordRepo,
	}
}

func (s *service) Create(user *domain.User) (*dto.ResponseUserForJWT, error) {
	// set userID
	user.ID = uuid.New().String()

	//hash user password
	pwd, err := s.passwordRepo.GenerateHashedPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = string(pwd)

	//set status to active
	user.Active = true

	var responseUserJWT dto.ResponseUserForJWT

	responseUserJWT.ID = user.ID
	responseUserJWT.Username = user.Username

	return &responseUserJWT, s.userRepo.Create(user)
}

func (s *service) GetByID(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.GetByID(id)
}

func (s *service) GetByUsername(username string) (*domain.User, error) {
	if username == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.GetByUsername(username)
}

func (s *service) UpdatePassword(id string, userPasswordDTO *dto.UserUpdatePasswordDTO) error {
	var user domain.User

	if id == "" || userPasswordDTO.Password == "" {
		return errors.New("input cannot be blank")
	}

	if err := controller.Validate.Struct(userPasswordDTO); err != nil {
		return err
	}

	user.ID = id

	pwd, err := s.passwordRepo.GenerateHashedPassword(userPasswordDTO.Password)
	if err != nil {
		return err
	}

	user.Password = string(pwd)

	return s.userRepo.UpdatePassword(id, &user)
}

func (s *service) AddAddress(id string, address []dto.AddressDTO) error {
	var user domain.User

	var addresses []domain.Address

	for i  := range address {
	var	_address domain.Address
		_address.Line1 = address[i].Line1
		_address.Line2 = address[i].Line2
		_address.Line3 = address[i].Line3
		_address.City = address[i].City
		_address.State = address[i].State
		_address.Country = address[i].Country

	addresses = append(addresses, _address)
	}

	user.Address = addresses

	return s.userRepo.AddAddress(id, &user)

}

func (s *service) Delete(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.Delete(id)
}
