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
