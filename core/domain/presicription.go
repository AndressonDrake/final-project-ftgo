package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type PrescriptionRepository interface {
	Get() (data []model.Prescription, err error)
	FindByID(id int) (data model.Prescription, err error)
}
type PrescriptionUsecase interface {
	Get() (data []model.Prescription, message, detail string, err error)
	GetByID(id int) (data model.Prescription, message, detail string, err error)
}
type PrescriptionHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}
