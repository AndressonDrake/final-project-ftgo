	package domain

	import (
		"core-healtcare.com/model"
		"gorm.io/gorm"
	)

	type AppointmentRepository interface {
		Begin() (tx *gorm.DB)
		Commit(tx *gorm.DB) (err error)
		Rollback(tx *gorm.DB) (err error)
		Create(tx *gorm.DB, request model.Appointment) (err error)
	}

	type AppointmentUsecase interface {
		Create(request model.CreateAppointment) (err error)
	}

	type AppointmentHandler interface {
	}