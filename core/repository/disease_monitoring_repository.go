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

func (r *diseaseMonitoringRepository) Get() (data []model.DiseaseMonitoring, err error) {
	err = r.db.Preload("ICD10").Find(&data).Error
	return
}
