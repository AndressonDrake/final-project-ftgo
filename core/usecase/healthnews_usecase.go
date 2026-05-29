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

func (u *healthNewsUsecase) Get() (data []model.HealthNews, message, detail string, err error) {
	data, err = u.healthNewsRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get health news"
	return
}

func (u *healthNewsUsecase) GetByID(id int) (data model.HealthNews, message, detail string, err error) {
	data, err = u.healthNewsRepository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get health news by ID"
	return
}
