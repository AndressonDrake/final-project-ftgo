package model

type ResponseSuccessGetPayment struct {
	Message string    `json:"message"`
	Data    []Payment `json:"data"`
}
