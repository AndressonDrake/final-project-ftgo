package model

type ResponseSuccessGetPayment struct {
	Message string    `json:"message"`
	Data    []Payment `json:"data"`
}

type ResponseSuccessGetPaymentByID struct {
	Message string  `json:"message"`
	Data    Payment `json:"data"`
}
