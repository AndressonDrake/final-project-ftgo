package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type PrescriptionRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Prescription) (err error)
}

type PrescriptionUsecase interface {
	Create(request model.CreatePrescription) (err error)
}

type PrescriptionHandler interface {
}