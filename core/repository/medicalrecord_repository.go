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

func (m *medicalRecordRepository) Begin() (tx *gorm.DB) {
	tx = m.db.Begin()
	return
}

func (m *medicalRecordRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (m *medicalRecordRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (m *medicalRecordRepository) Create(tx *gorm.DB, request model.MedicalRecord) (err error) {
	err = tx.Create(&request).Error
	return
}

func (r *medicalRecordRepository) Get() (data []model.MedicalRecord, err error) {
	err = r.db.Preload("Appointment").Preload("ICD10").Find(&data).Error
	return
}

func (r *medicalRecordRepository) FindByID(id int) (data model.MedicalRecord, err error) {
	err = r.db.Preload("Appointment").Preload("ICD10").First(&data, id).Error
	return
}