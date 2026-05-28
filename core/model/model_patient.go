package model

type ResponseSuccessGetPatient struct {
	Message string    `json:"message"`
	Data    []Patient `json:"data"`
}
