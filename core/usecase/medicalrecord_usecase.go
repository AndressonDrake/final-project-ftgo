package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type medicalRecordUsecase struct {
	medicalRecordRepository domain.MedicalRecordRepository
}

func MedicalRecordUsecase(
	medicalRecordRepository domain.MedicalRecordRepository,
) domain.MedicalRecordUsecase {
	return &medicalRecordUsecase{
		medicalRecordRepository: medicalRecordRepository,
	}
}

func (m *medicalRecordUsecase) Create(
	request model.CreateMedicalRecord,
) (err error) {

	var req model.MedicalRecord

	req.IDAppointment = request.IdAppointment
	req.IDICD = request.IdIcd
	req.HasilLab = request.HasilLab
	req.HasilRadiologi = request.HasilRadiologi
	req.Tindakan = request.Tindakan
	req.Catatan = request.Catatan

	tx := m.medicalRecordRepository.Begin()

	err = m.medicalRecordRepository.Create(tx, req)
	if err != nil {
		m.medicalRecordRepository.Rollback(tx)
		return
	}

	m.medicalRecordRepository.Commit(tx)

	return
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

