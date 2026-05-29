package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type PatientRepository interface {
	Get() (data []model.Patient, err error)
	FindByID(id int) (data model.Patient, err error)
}
type PatientUsecase interface {
	Get() (data []model.Patient, message, detail string, err error)
	GetByID(id int) (data model.Patient, message, detail string, err error)
}
type PatientHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}
