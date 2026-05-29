package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type PatientStatusRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.PatientStatus) (err error)
}

type PatientStatusUsecase interface {
	Create(request model.CreatePatientStatus) (err error)
}

type PatientStatusHandler interface {
}