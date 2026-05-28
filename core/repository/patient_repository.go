package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type patientRepository struct {
	db *gorm.DB
}

func PatientRepository(db *gorm.DB) domain.PatientRepository {
	return &patientRepository{db: db}
}

func (r *patientRepository) Get() (data []model.Patient, err error) {
	err = r.db.Preload("Status").Find(&data).Error
	return
}
