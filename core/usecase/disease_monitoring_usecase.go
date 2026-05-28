package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type diseaseMonitoringUsecase struct {
	diseaseMonitoringRepository domain.DiseaseMonitoringRepository
}

func DiseaseMonitoringUsecase(diseaseMonitoringRepository domain.DiseaseMonitoringRepository) domain.DiseaseMonitoringUsecase {
	return &diseaseMonitoringUsecase{diseaseMonitoringRepository: diseaseMonitoringRepository}
}

func (u *diseaseMonitoringUsecase) Get() (data []model.DiseaseMonitoring, message, detail string, err error) {
	data, err = u.diseaseMonitoringRepository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get disease monitoring"
	return
}
