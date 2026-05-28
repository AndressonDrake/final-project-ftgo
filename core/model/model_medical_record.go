package model

type ResponseSuccessGetMedicalRecord struct {
	Message string          `json:"message"`
	Data    []MedicalRecord `json:"data"`
}
