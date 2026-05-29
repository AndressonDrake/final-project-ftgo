package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type diseaseMonitoringRepository struct {
	db *gorm.DB
}

func DiseaseMonitoringRepository(db *gorm.DB) domain.DiseaseMonitoringRepository {
	return &diseaseMonitoringRepository{db: db}
}

func (d *diseaseMonitoringRepository) Begin() (tx *gorm.DB) {
	tx = d.db.Begin()
	return
}

func (d *diseaseMonitoringRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (d *diseaseMonitoringRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (d *diseaseMonitoringRepository) Create(tx *gorm.DB, request model.DiseaseMonitoring) (err error) {
	err = tx.Create(&request).Error
	return
}


func (r *diseaseMonitoringRepository) Get() (data []model.DiseaseMonitoring, err error) {
	err = r.db.Preload("ICD10").Find(&data).Error
	return
}

func (r *diseaseMonitoringRepository) FindByID(id int) (data model.DiseaseMonitoring, err error) {
	err = r.db.Preload("ICD10").First(&data, id).Error
	return
}
