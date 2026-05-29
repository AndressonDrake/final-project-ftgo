package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type healthNewsUsecase struct {
	healthNewsRepository domain.HealthNewsRepository
}

func HealthNewsUsecase(
	healthNewsRepository domain.HealthNewsRepository,
) domain.HealthNewsUsecase {
	return &healthNewsUsecase{
		healthNewsRepository: healthNewsRepository,
	}
}

func (h *healthNewsUsecase) Create(
	request model.CreateHealthNews,
) (err error) {

	var req model.HealthNews

	req.Judul = request.Judul
	req.Sumber = request.Sumber
	req.Kategori = request.Kategori
	req.URL = request.Url

	tx := h.healthNewsRepository.Begin()

	err = h.healthNewsRepository.Create(tx, req)
	if err != nil {
		h.healthNewsRepository.Rollback(tx)
		return
	}

	h.healthNewsRepository.Commit(tx)

	return
}