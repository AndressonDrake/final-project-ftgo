package model

type ResponseSuccessGetPayment struct {
	Message string    `json:"message"`
	Data    []Payment `json:"data"`
}

type ResponseSuccessGetPaymentByID struct {
	Message string  `json:"message"`
	Data    Payment `json:"data"`
}
type CreatePayment struct {
	IdAppointment    int     `json:"id_appointment"`
	Total            float64 `json:"total"`
	MetodePembayaran string  `json:"metode_pembayaran"`
	StatusPembayaran string  `json:"status_pembayaran"`
}
