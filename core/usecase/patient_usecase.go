package usecase

import (
	"time"

	"core-healtcare.com/domain"
	"core-healtcare.com/helper"
	"core-healtcare.com/model"
)

type patientUsecase struct {
	patientRepository domain.PatientRepository
}

func PatientUsecase(patientRepository domain.PatientRepository) domain.PatientUsecase {
	return &patientUsecase{patientRepository: patientRepository}
}

func (p *patientUsecase) Create(request model.CreatePatient) (err error) {

	req := model.Patient{
		IDStatus:      request.IdStatus,
		Nama:          request.Nama,
		NIK:           request.Nik,
		Gender:        request.Gender,
		Alamat:        request.Alamat,
		NoHP:          request.NoHp,
		GolonganDarah: request.GolonganDarah,
	}

	req.TanggalLahir, _ = time.Parse(helper.DATELAYOUT, request.TanggalLahir)

	tx := p.patientRepository.Begin()

	err = p.patientRepository.Create(tx, req)
	if err != nil {
		p.patientRepository.Rollback(tx)
		return
	}

	p.patientRepository.Commit(tx)

	return
}

func (u *patientUsecase) Get() (data []model.Patient, message, detail string, err error) {
	data, err = u.patientRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get patient"
	return
}

func (u *patientUsecase) GetByID(id int) (data model.Patient, message, detail string, err error) {
	data, err = u.patientRepository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get patient by ID"
	return
}