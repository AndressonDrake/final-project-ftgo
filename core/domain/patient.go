package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type PatientRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Patient) (err error)
}

type PatientUsecase interface {
	Create(request model.CreatePatient) (err error)
}

type PatientHandler interface {
}