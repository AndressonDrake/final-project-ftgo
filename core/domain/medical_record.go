package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MedicalRecordRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
	Create(tx *gorm.DB, request model.MedicalRecord) (err error)
	Get() (data []model.MedicalRecord, err error)
}

type MedicalRecordUsecase interface {
	Create(request model.CreateMedicalRecord) (err error)
	Get() (data []model.MedicalRecord, message, detail string, err error)
}

type MedicalRecordHandler interface {
	Get(c echo.Context) (err error)
	Create(c echo.Context) (err error)
	Update(c echo.Context) (err error)
	Delete(c echo.Context) (err error)
}
