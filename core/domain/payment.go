package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Payment) (err error)
}

type PaymentUsecase interface {
	Create(request model.CreatePayment) (err error)
}

type PaymentHandler interface {
}