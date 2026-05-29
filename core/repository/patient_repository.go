// ======================================================
// PATIENT REPOSITORY
// ======================================================

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

func (p *patientRepository) Begin() (tx *gorm.DB) {
	tx = p.db.Begin()
	return
}

func (p *patientRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (p *patientRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (p *patientRepository) Create(tx *gorm.DB, request model.Patient) (err error) {
	err = tx.Create(&request).Error
	return
}

func (r *patientRepository) Get() (data []model.Patient, err error) {
	err = r.db.Preload("Status").Find(&data).Error
	return
}

func (r *patientRepository) FindByID(id int) (data model.Patient, err error) {
	err = r.db.Preload("Status").First(&data, id).Error
	return
}