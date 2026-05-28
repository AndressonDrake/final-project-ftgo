package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type patientUsecase struct {
	patientRepository domain.PatientRepository
}

func PatientUsecase(patientRepository domain.PatientRepository) domain.PatientUsecase {
	return &patientUsecase{patientRepository: patientRepository}
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
