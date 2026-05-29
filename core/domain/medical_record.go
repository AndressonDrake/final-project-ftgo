package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type MedicalRecordRepository interface {
	Get() (data []model.MedicalRecord, err error)
	FindByID(id int) (data model.MedicalRecord, err error)
}
type MedicalRecordUsecase interface {
	Get() (data []model.MedicalRecord, message, detail string, err error)
	GetByID(id int) (data model.MedicalRecord, message, detail string, err error)
}
type MedicalRecordHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}
