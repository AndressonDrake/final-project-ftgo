package model

type CreateMedicalRecord struct {
	IDAppointment  int    `json:"id_appointment"`
	IDICD          int    `json:"id_icd"`
	HasilLab       string `json:"hasil_lab"`
	HasilRadiologi string `json:"hasil_radiologi"`
	Tindakan       string `json:"tindakan"`
	Catatan        string `json:"catatan"`
}

type ResponseSuccessGetMedicalRecord struct {
	Message string          `json:"message"`
	Data    []MedicalRecord `json:"data"`
}
