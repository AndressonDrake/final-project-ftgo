package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type HealthNewsRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.HealthNews) (err error)
}

type HealthNewsUsecase interface {
	Create(request model.CreateHealthNews) (err error)
}

type HealthNewsHandler interface {
}