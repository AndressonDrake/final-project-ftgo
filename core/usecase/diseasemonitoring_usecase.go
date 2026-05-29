package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type diseaseMonitoringUsecase struct {
	diseaseMonitoringRepository domain.DiseaseMonitoringRepository
}

func DiseaseMonitoringUsecase(
	diseaseMonitoringRepository domain.DiseaseMonitoringRepository,
) domain.DiseaseMonitoringUsecase {
	return &diseaseMonitoringUsecase{
		diseaseMonitoringRepository: diseaseMonitoringRepository,
	}
}

func (d *diseaseMonitoringUsecase) Create(
	request model.CreateDiseaseMonitoring,
) (err error) {

	var req model.DiseaseMonitoring

	req.IDICD = request.IdIcd
	req.Negara = request.Negara
	req.TotalKasus = request.TotalKasus
	req.TotalKematian = request.TotalKematian
	req.TotalSembuh = request.TotalSembuh

	tx := d.diseaseMonitoringRepository.Begin()

	err = d.diseaseMonitoringRepository.Create(tx, req)
	if err != nil {
		d.diseaseMonitoringRepository.Rollback(tx)
		return
	}

	d.diseaseMonitoringRepository.Commit(tx)

	return
}