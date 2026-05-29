package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type DiseaseMonitoringRepository interface {
	Get() (data []model.DiseaseMonitoring, err error)
	FindByID(id int) (data model.DiseaseMonitoring, err error)
}
type DiseaseMonitoringUsecase interface {
	Get() (data []model.DiseaseMonitoring, message, detail string, err error)
	GetByID(id int) (data model.DiseaseMonitoring, message, detail string, err error)
}
type DiseaseMonitoringHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}
