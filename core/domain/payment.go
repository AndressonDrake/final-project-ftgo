package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type PaymentRepository interface {
	Get() (data []model.Payment, err error)
	FindByID(id int) (data model.Payment, err error)
}
type PaymentUsecase interface {
	Get() (data []model.Payment, message, detail string, err error)
	GetByID(id int) (data model.Payment, message, detail string, err error)
}
type PaymentHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}
