package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AppointmentRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
	Create(tx *gorm.DB, request model.Appointment) (err error)
	Get() (data []model.Appointment, err error)
	FindByID(id int) (data model.Appointment, err error) // 🆕
}
type AppointmentUsecase interface {
	Create(request model.CreateAppointment) (err error)
	Get() (data []model.Appointment, message, detail string, err error)
	GetByID(id int) (data model.Appointment, message, detail string, err error) // 🆕
}
type AppointmentHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error) // 🆕
}
