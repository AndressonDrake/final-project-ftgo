package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PrescriptionRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Prescription) (err error)
	Get() (data []model.Prescription, err error)
	FindByID(id int) (data model.Prescription, err error)
}

type PrescriptionUsecase interface {
	Create(request model.CreatePrescription) (err error)
	Get() (data []model.Prescription, message, detail string, err error)
	GetByID(id int) (data model.Prescription, message, detail string, err error)
}

type PrescriptionHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}