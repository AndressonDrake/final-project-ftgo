package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type patientStatusUsecase struct {
	patientStatusRepository domain.PatientStatusRepository
}

func PatientStatusUsecase(patientStatusRepository domain.PatientStatusRepository) domain.PatientStatusUsecase {
	return &patientStatusUsecase{patientStatusRepository: patientStatusRepository}
}

func (p *patientStatusUsecase) Create(request model.CreatePatientStatus) (err error) {
	var req model.PatientStatus
	req.NamaStatus = request.NamaStatus

	tx := p.patientStatusRepository.Begin()
	err = p.patientStatusRepository.Create(tx, req)
	if err != nil {
		p.patientStatusRepository.Rollback(tx)
		return
	}
	p.patientStatusRepository.Commit(tx)
	return
}