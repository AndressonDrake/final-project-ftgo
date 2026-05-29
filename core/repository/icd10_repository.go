package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type icd10Repository struct {
	db *gorm.DB
}

func ICD10Repository(db *gorm.DB) domain.ICD10Repository {
	return &icd10Repository{db: db}
}

func (i *icd10Repository) Begin() (tx *gorm.DB) {
	tx = i.db.Begin()
	return
}

func (i *icd10Repository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (i *icd10Repository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (i *icd10Repository) Create(tx *gorm.DB, request model.ICD10) (err error) {
	err = tx.Create(&request).Error
	return
}

func (r *icd10Repository) Get() (data []model.ICD10, err error) {
	err = r.db.Find(&data).Error
	return
}

func (r *icd10Repository) FindByID(id int) (data model.ICD10, err error) {
	err = r.db.First(&data, id).Error
	return
}