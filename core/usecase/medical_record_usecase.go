package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type medicalRecordUsecase struct {
	medicalRecordRepository domain.MedicalRecordRepository
}

func MedicalRecordUsecase(medicalRecordRepository domain.MedicalRecordRepository) domain.MedicalRecordUsecase {
	return &medicalRecordUsecase{medicalRecordRepository: medicalRecordRepository}
}

func (u *medicalRecordUsecase) Get() (data []model.MedicalRecord, message, detail string, err error) {
	data, err = u.medicalRecordRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get medical record"
	return
}

func (u *medicalRecordUsecase) GetByID(id int) (data model.MedicalRecord, message, detail string, err error) {
	data, err = u.medicalRecordRepository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get medical record by ID"
	return
}
