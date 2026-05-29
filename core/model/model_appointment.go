package model

type ResponseSuccessGetAppointment struct {
	Message string        `json:"message"`
	Data    []Appointment `json:"data"`
}

type ResponseSuccessGetAppointmentByID struct {
	Message string      `json:"message"`
	Data    Appointment `json:"data"`
}
