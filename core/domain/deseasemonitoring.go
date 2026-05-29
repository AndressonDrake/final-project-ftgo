package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DiseaseMonitoringRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.DiseaseMonitoring) (err error)
	Get() (data []model.DiseaseMonitoring, err error)
	FindByID(id int) (data model.DiseaseMonitoring, err error)
}

type DiseaseMonitoringUsecase interface {
	Create(request model.CreateDiseaseMonitoring) (err error)
	Get() (data []model.DiseaseMonitoring, message, detail string, err error)
	GetByID(id int) (data model.DiseaseMonitoring, message, detail string, err error)
}

type DiseaseMonitoringHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}