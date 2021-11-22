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
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	UpdatePassword(id string, userPasswordDTO *dto.UserUpdatePasswordDTO) error
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

func (s *service) Create(user *domain.User) error {
	// set userID
	user.ID = uuid.New().String()

	//hash user password
	pwd, err := s.passwordRepo.GenerateHashedPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(pwd)

	//set status to active
	user.Active = true

	return s.userRepo.Create(user)
}

func (s *service) GetByID(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.GetByID(id)
}

func (s *service) GetByUsername(username string) (*domain.User, error){
	if username == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.GetByUsername(username)
}

func (s *service) UpdatePassword(id string, userPasswordDTO *dto.UserUpdatePasswordDTO) error {
	var user domain.User 

	if id == ""  || userPasswordDTO.Password == "" {
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

func (s *service) Delete(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.Delete(id)
}
