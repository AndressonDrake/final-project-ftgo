package usecase

import (
	"time"

	"core-healtcare.com/domain"
	"core-healtcare.com/helper"
	"core-healtcare.com/model"
)

type medicineUsecase struct {
	medicineRepository domain.MedicineRepository
}

func MedicineUsecase(medicineRepository domain.MedicineRepository) domain.MedicineUsecase {
	return &medicineUsecase{medicineRepository: medicineRepository}
}

func (m *medicineUsecase) Create(request model.CreateMedicine) (err error) {
	var req model.Medicine

	req.NamaObat = request.NamaObat
	req.Kategori = request.Kategori
	req.Harga = request.Harga
	req.ExpiredDate, _ = time.Parse(helper.DATELAYOUT, request.ExpiredDate)
	req.Stok = request.Stok

	tx := m.medicineRepository.Begin()

	err = m.medicineRepository.Create(tx, req)
	if err != nil {
		m.medicineRepository.Rollback(tx)
		return
	}

	m.medicineRepository.Commit(tx)

	return
}
