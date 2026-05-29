package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Payment) (err error)
	Get() (data []model.Payment, err error)
	FindByID(id int) (data model.Payment, err error)
}

type PaymentUsecase interface {
	Create(request model.CreatePayment) (err error)
	Get() (data []model.Payment, message, detail string, err error)
	GetByID(id int) (data model.Payment, message, detail string, err error)
}

type PaymentHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}