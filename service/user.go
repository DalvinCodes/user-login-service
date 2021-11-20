package service

import (
	"errors"
	"github.com/google/uuid"
	"user-login-service/domain"
	"user-login-service/pkg"
	"user-login-service/repository"
)

type UserService interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	Update(id string, user *domain.User) error
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

func (s *service) Update(id string, user *domain.User) error {
	return s.userRepo.Update(id, user)
}

func (s *service) Delete(id string) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("input cannot be blank")
	}
	return s.userRepo.Delete(id)
}
