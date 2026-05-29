package model

type ResponseSuccessGetDiseaseMonitoring struct {
	Message string              `json:"message"`
	Data    []DiseaseMonitoring `json:"data"`
}

type ResponseSuccessGetDiseaseMonitoringByID struct {
	Message string            `json:"message"`
	Data    DiseaseMonitoring `json:"data"`
}

type CreateDiseaseMonitoring struct {
	IdIcd         int    `json:"id_icd"`
	Negara        string `json:"negara"`
	TotalKasus    int    `json:"total_kasus"`
	TotalKematian int    `json:"total_kematian"`
	TotalSembuh   int    `json:"total_sembuh"`
	TanggalUpdate string `json:"tanggal_update"`
}