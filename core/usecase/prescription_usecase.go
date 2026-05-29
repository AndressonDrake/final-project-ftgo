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