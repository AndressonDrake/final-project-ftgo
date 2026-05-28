package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type prescriptionRepository struct {
	db *gorm.DB
}

func PrescriptionRepository(db *gorm.DB) domain.PrescriptionRepository {
	return &prescriptionRepository{db: db}
}

func (r *prescriptionRepository) Get() (data []model.Prescription, err error) {
	err = r.db.Preload("Medicine").Preload("MedicalRecord").Find(&data).Error
	return
}
