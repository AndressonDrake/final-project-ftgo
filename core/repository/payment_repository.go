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

func (p *paymentRepository) Begin() (tx *gorm.DB) {
	tx = p.db.Begin()
	return
}

func (p *paymentRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (p *paymentRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (p *paymentRepository) Create(tx *gorm.DB, request model.Payment) (err error) {
	err = tx.Create(&request).Error
	return
}

func (r *paymentRepository) Get() (data []model.Payment, err error) {
	err = r.db.Preload("Appointment").Find(&data).Error
	return
}

func (r *paymentRepository) FindByID(id int) (data model.Payment, err error) {
	err = r.db.Preload("Appointment").First(&data, id).Error
	return
}