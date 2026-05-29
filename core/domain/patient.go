package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PatientRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Patient) (err error)
	Get() (data []model.Patient, err error)
	FindByID(id int) (data model.Patient, err error)
}

type PatientUsecase interface {
	Create(request model.CreatePatient) (err error)
	Get() (data []model.Patient, message, detail string, err error)
	GetByID(id int) (data model.Patient, message, detail string, err error)
}

type PatientHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}