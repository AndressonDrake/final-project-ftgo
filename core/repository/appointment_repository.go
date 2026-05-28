package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type appointmentRepository struct {
	db *gorm.DB
}

func AppointmentRepository(db *gorm.DB) domain.AppointmentRepository {
	return &appointmentRepository{db: db}
}

func (ar *appointmentRepository) Begin() (tx *gorm.DB) {
	tx = ar.db.Begin()
	return
}

func (ar *appointmentRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (ar *appointmentRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (ar *appointmentRepository) Create(tx *gorm.DB, request model.Appointment) (err error) {
	err = tx.Create(&request).Error
	return
}

func (ar *appointmentRepository) Get() (data []model.Appointment, err error) {
	err = ar.db.Preload("Patient").Preload("Doctor").Find(&data).Error
	return
}
