package repository

import (
	"user-login-service/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByUsername(username string) (*domain.User, error)
	UpdatePassword(id string, user *domain.User) error
	AddAddress(id string, user *domain.User) error
	Delete(id string) (*domain.User, error)
}

type repo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *repo) GetByID(id string) (*domain.User, error) {
	var user *domain.User
	if err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) GetByUsername(username string) (*domain.User, error) {
	var user *domain.User

	if err := r.db.Preload(clause.Associations).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repo) UpdatePassword(id string, user *domain.User) error {
	return r.db.Model(user).Select("password").Where("id = ?", id).First(user).Updates(user).Error
}

func (r *repo) AddAddress(id string, userDTO *domain.User) error {
	var user *domain.User
	
	if err := r.db.Preload(clause.Associations).Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	return r.db.Model(&user).Association("Address").Append(userDTO.Address)
}

func (r *repo) Delete(id string) (*domain.User, error) {
	var user *domain.User
	if err := r.db.Preload(clause.Associations).Where("id = ?", id).Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
