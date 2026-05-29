package model

type CreatePayment struct {
	IdAppointment     int     `json:"id_appointment"`
	Total             float64 `json:"total"`
	MetodePembayaran  string  `json:"metode_pembayaran"`
	StatusPembayaran  string  `json:"status_pembayaran"`
}