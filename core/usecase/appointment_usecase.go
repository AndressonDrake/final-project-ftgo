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

func (au *appointmentUsecase) Create(request model.CreateAppointment) (err error) {
	var req model.Appointment

	req.IDPatient = request.IDPatient
	req.IDDoctor = request.IDDoctor
	req.Tanggal, _ = time.Parse(helper.DATELAYOUT, request.Tanggal)
	req.Keluhan = request.Keluhan
	req.TekananDarah = request.TekananDarah
	req.SuhuTubuh = request.SuhuTubuh
	req.BeratBadan = request.BeratBadan
	req.Status = request.Status

	tx := au.appointmentRepository.Begin()

	err = au.appointmentRepository.Create(tx, req)
	if err != nil {
		au.appointmentRepository.Rollback(tx)
		return
	}

	au.appointmentRepository.Commit(tx)

	return
}

func (au *appointmentUsecase) Get() (data []model.Appointment, message, detail string, err error) {
	data, err = au.appointmentRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}

	message = "success get appointment"

	return
}
