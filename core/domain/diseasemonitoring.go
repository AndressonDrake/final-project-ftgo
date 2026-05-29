package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type DiseaseMonitoringRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.DiseaseMonitoring) (err error)
}

type DiseaseMonitoringUsecase interface {
	Create(request model.CreateDiseaseMonitoring) (err error)
}

type DiseaseMonitoringHandler interface {
}