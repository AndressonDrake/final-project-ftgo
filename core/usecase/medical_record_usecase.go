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

func (mru *medicalRecordUsecase) Create(request model.CreateMedicalRecord) (err error) {
	var req model.MedicalRecord

	req.IDAppointment = request.IDAppointment
	req.IDICD = request.IDICD
	req.HasilLab = request.HasilLab
	req.HasilRadiologi = request.HasilRadiologi
	req.Tindakan = request.Tindakan
	req.Catatan = request.Catatan

	tx := mru.medicalRecordRepository.Begin()

	err = mru.medicalRecordRepository.Create(tx, req)
	if err != nil {
		mru.medicalRecordRepository.Rollback(tx)
		return
	}

	mru.medicalRecordRepository.Commit(tx)

	return
}

func (mru *medicalRecordUsecase) Get() (data []model.MedicalRecord, message, detail string, err error) {
	data, err = mru.medicalRecordRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}

	message = "success get medical record"

	return
}
