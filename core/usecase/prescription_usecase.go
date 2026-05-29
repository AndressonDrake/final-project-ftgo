package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type prescriptionUsecase struct {
	prescriptionRepository domain.PrescriptionRepository
}

func PrescriptionUsecase(prescriptionRepository domain.PrescriptionRepository) domain.PrescriptionUsecase {
	return &prescriptionUsecase{prescriptionRepository: prescriptionRepository}
}

func (u *prescriptionUsecase) Get() (data []model.Prescription, message, detail string, err error) {
	data, err = u.prescriptionRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get prescription"
	return
}

func (u *prescriptionUsecase) GetByID(id int) (data model.Prescription, message, detail string, err error) {
	data, err = u.prescriptionRepository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get prescription by ID"
	return
}
