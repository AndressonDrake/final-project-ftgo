package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type medicalRecordRepository struct {
	db *gorm.DB
}

func MedicalRecordRepository(db *gorm.DB) domain.MedicalRecordRepository {
	return &medicalRecordRepository{db: db}
}

func (r *medicalRecordRepository) Get() (data []model.MedicalRecord, err error) {
	err = r.db.Preload("Appointment").Preload("ICD10").Find(&data).Error
	return
}
