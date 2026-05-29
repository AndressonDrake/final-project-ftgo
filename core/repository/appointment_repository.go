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

func (r *appointmentRepository) Get() (data []model.Appointment, err error) {
	err = r.db.Preload("Patient").Preload("Doctor").Find(&data).Error
	return
}

func (r *appointmentRepository) FindByID(id int) (data model.Appointment, err error) {
	err = r.db.Preload("Patient").Preload("Doctor").First(&data, id).Error
	return
}
