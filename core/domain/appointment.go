package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type AppointmentRepository interface {
	Get() (data []model.Appointment, err error)
	FindByID(id int) (data model.Appointment, err error) // 🆕
}
type AppointmentUsecase interface {
	Get() (data []model.Appointment, message, detail string, err error)
	GetByID(id int) (data model.Appointment, message, detail string, err error) // 🆕
}
type AppointmentHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error) // 🆕
}
