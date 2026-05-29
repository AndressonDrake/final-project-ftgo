package domain

import (
	"gorm.io/gorm"
	"core-healtcare.com/model"
)

type UserRepository interface {
	Begin() *gorm.DB
	Commit(tx *gorm.DB) error
	Rollback(tx *gorm.DB) error

	Create(tx *gorm.DB, request model.User) error
	FindAll() ([]model.User, error)
	FindById(id int) (model.User, error)
	Update(tx *gorm.DB, request model.User) error
	Delete(tx *gorm.DB, id int) error
}

type UserUsecase interface {
	Create(request model.CreateUser) (err error)
}

type UserHandler interface {
}	