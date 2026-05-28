package model

type CreateAppointment struct {
	IDPatient    int     `json:"id_patient"`
	IDDoctor     int     `json:"id_doctor"`
	Tanggal      string  `json:"tanggal"`
	Keluhan      string  `json:"keluhan"`
	TekananDarah string  `json:"tekanan_darah"`
	SuhuTubuh    float64 `json:"suhu_tubuh"`
	BeratBadan   float64 `json:"berat_badan"`
	Status       string  `json:"status"`
}

type ResponseSuccessGetAppointment struct {
	Message string        `json:"message"`
	Data    []Appointment `json:"data"`
}
