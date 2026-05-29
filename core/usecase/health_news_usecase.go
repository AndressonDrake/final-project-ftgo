package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type healthNewsUsecase struct {
	healthNewsRepository domain.HealthNewsRepository
}

func HealthNewsUsecase(healthNewsRepository domain.HealthNewsRepository) domain.HealthNewsUsecase {
	return &healthNewsUsecase{healthNewsRepository: healthNewsRepository}
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
