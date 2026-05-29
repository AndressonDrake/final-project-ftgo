package usecase

import (
	"time"

	"core-healtcare.com/domain"
	"core-healtcare.com/helper"
	"core-healtcare.com/model"
)

type appointmentUsecase struct {
	appointmentRepository domain.AppointmentRepository
}


func AppointmentUsecase(appointmentRepository domain.AppointmentRepository) domain.AppointmentUsecase {
	return &appointmentUsecase{appointmentRepository: appointmentRepository}
}

func (a *appointmentUsecase) Create(request model.CreateAppointment) (err error) {

	req := model.Appointment{
		IDPatient:    request.IdPatient,
		IDDoctor:     request.IdDoctor,
		Keluhan:      request.Keluhan,
		TekananDarah: request.TekananDarah,
		SuhuTubuh:    request.SuhuTubuh,
		BeratBadan:   request.BeratBadan,
		Status:       request.Status,
	}

	req.Tanggal, _ = time.Parse(helper.DATELAYOUT, request.Tanggal)

	tx := a.appointmentRepository.Begin()

	err = a.appointmentRepository.Create(tx, req)
	if err != nil {
		a.appointmentRepository.Rollback(tx)
		return
	}

	a.appointmentRepository.Commit(tx)

	return
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