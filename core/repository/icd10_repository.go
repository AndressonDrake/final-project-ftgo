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

func (r *icd10Repository) Get() (data []model.ICD10, err error) {
	err = r.db.Find(&data).Error
	return
}

func (r *icd10Repository) FindByID(id int) (data model.ICD10, err error) {
	err = r.db.First(&data, id).Error
	return
}
