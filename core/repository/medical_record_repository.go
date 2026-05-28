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

func (mrr *medicalRecordRepository) Begin() (tx *gorm.DB) {
	tx = mrr.db.Begin()
	return
}

func (mrr *medicalRecordRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (mrr *medicalRecordRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (mrr *medicalRecordRepository) Create(tx *gorm.DB, request model.MedicalRecord) (err error) {
	err = tx.Create(&request).Error
	return
}

func (mrr *medicalRecordRepository) Get() (data []model.MedicalRecord, err error) {
	err = mrr.db.Preload("Appointment").Preload("ICD10").Find(&data).Error
	return
}
