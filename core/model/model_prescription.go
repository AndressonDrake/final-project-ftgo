package model

type ResponseSuccessGetPrescription struct {
	Message string         `json:"message"`
	Data    []Prescription `json:"data"`
}

type ResponseSuccessGetPrescriptionByID struct {
	Message string       `json:"message"`
	Data    Prescription `json:"data"`
}
