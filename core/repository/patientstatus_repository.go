package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type patientStatusRepository struct {
	db *gorm.DB
}

func PatientStatusRepository(db *gorm.DB) domain.PatientStatusRepository {
	return &patientStatusRepository{db: db}
}

func (p *patientStatusRepository) Begin() (tx *gorm.DB) {
	tx = p.db.Begin()
	return
}

func (p *patientStatusRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (p *patientStatusRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (p *patientStatusRepository) Create(tx *gorm.DB, request model.PatientStatus) (err error) {
	err = tx.Create(&request).Error
	return
}