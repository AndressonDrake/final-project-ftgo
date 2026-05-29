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

func (p *prescriptionRepository) Begin() (tx *gorm.DB) {
	tx = p.db.Begin()
	return
}

func (p *prescriptionRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (p *prescriptionRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (p *prescriptionRepository) Create(tx *gorm.DB, request model.Prescription) (err error) {
	err = tx.Create(&request).Error
	return
}
func (r *prescriptionRepository) Get() (data []model.Prescription, err error) {
	err = r.db.Preload("Medicine").Preload("MedicalRecord").Find(&data).Error
	return
}

func (r *prescriptionRepository) FindByID(id int) (data model.Prescription, err error) {
	err = r.db.Preload("Medicine").Preload("MedicalRecord").First(&data, id).Error
	return
}