package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type prescriptionUsecase struct {
	prescriptionRepository domain.PrescriptionRepository
}

func PrescriptionUsecase(repo domain.PrescriptionRepository) domain.PrescriptionUsecase {
	return &prescriptionUsecase{prescriptionRepository: repo}
}

func (p *prescriptionUsecase) Create(request model.CreatePrescription) (err error) {

	req := model.Prescription{
		IDRecord:    request.IdRecord,
		IDObat:      request.IdObat,
		Jumlah:      request.Jumlah,
		AturanPakai: request.AturanPakai,
	}

	tx := p.prescriptionRepository.Begin()

	err = p.prescriptionRepository.Create(tx, req)
	if err != nil {
		p.prescriptionRepository.Rollback(tx)
		return
	}

	p.prescriptionRepository.Commit(tx)

	return
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