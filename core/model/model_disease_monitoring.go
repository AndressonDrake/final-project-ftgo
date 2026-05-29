package model

type ResponseSuccessGetDiseaseMonitoring struct {
	Message string              `json:"message"`
	Data    []DiseaseMonitoring `json:"data"`
}

type ResponseSuccessGetDiseaseMonitoringByID struct {
	Message string            `json:"message"`
	Data    DiseaseMonitoring `json:"data"`
}
