package model

type RequestGeneral struct {
	TrxType string      `json:"trx_type"`
	SubType string      `json:"sub_type"`
	Data    interface{} `json:"data"`
}

type ResponseSuccessPost struct {
	Message string `json:"message"`
}

type ResponseErrorPost struct {
	Detail  string `json:"detail"`
	Message string `json:"message"`
}
