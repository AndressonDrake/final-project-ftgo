package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Role) (err error)
}

type RoleUsecase interface {
	Create(request model.CreateRole) (err error)
}

type RoleHandler interface{}