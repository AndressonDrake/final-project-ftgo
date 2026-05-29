package repository

import (
    "fmt"

    "core-healtcare.com/domain"
    "core-healtcare.com/model"
    "gorm.io/gorm"
)

type appointmentRepository struct{
	db *gorm.DB
}

func AppointmentRepository(db *gorm.DB) domain.AppointmentRepository {
	return &appointmentRepository{
		db: db,
	}
}

func (a *appointmentRepository) Begin() (tx*gorm.DB) {
	tx = a.db.Begin()
	return
}

func (a *appointmentRepository) Commit(tx*gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (a *appointmentRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}


func (a *appointmentRepository) Create(tx *gorm.DB, request model.Appointment) (err error) {
    err = tx.Create(&request).Error
    fmt.Println("CREATE ERROR:", err)
    return
}