package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type paymentRepository struct {
	db *gorm.DB
}

func PaymentRepository(db *gorm.DB) domain.PaymentRepository {
	return &paymentRepository{db: db}
}

func (r *paymentRepository) Get() (data []model.Payment, err error) {
	err = r.db.Preload("Appointment").Find(&data).Error
	return
}
