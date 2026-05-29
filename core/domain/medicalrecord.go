package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type MedicalRecordRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.MedicalRecord) (err error)
}

type MedicalRecordUsecase interface {
	Create(request model.CreateMedicalRecord) (err error)
}

type MedicalRecordHandler interface {
}