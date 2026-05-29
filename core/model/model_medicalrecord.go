package model

type ResponseSuccessGetMedicalRecord struct {
	Message string          `json:"message"`
	Data    []MedicalRecord `json:"data"`
}

type ResponseSuccessGetMedicalRecordByID struct {
	Message string        `json:"message"`
	Data    MedicalRecord `json:"data"`
}
type CreateMedicalRecord struct {
	IdAppointment  int    `json:"id_appointment"`
	IdIcd          int    `json:"id_icd"`
	HasilLab       string `json:"hasil_lab"`
	HasilRadiologi string `json:"hasil_radiologi"`
	Tindakan       string `json:"tindakan"`
	Catatan        string `json:"catatan"`
}