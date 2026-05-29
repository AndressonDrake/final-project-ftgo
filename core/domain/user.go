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
	
}

type UserUsecase interface {
	Create(request model.CreateUser) (err error)
}

type UserHandler interface {
	
}	