package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type appointmentUsecase struct {
	appointmentRepository domain.AppointmentRepository
}

func AppointmentUsecase(appointmentRepository domain.AppointmentRepository) domain.AppointmentUsecase {
	return &appointmentUsecase{appointmentRepository: appointmentRepository}
}

func (u *appointmentUsecase) Get() (data []model.Appointment, message, detail string, err error) {
	data, err = u.appointmentRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get appointment"
	return
}

func (u *appointmentUsecase) GetByID(id int) (data model.Appointment, message, detail string, err error) {
	data, err = u.appointmentRepository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get appointment by ID"
	return
}
