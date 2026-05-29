package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
	Create(tx *gorm.DB, request model.Medicine) (err error)
}

type MedicineUsecase interface {
	Create(request model.CreateMedicine) (err error)
}

type MedicineHandler interface {
}	
