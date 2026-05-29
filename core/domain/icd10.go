package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type ICD10Repository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.ICD10) (err error)
}

type ICD10Usecase interface {
	Create(request model.CreateICD10) (err error)
}

type ICD10Handler interface {
}