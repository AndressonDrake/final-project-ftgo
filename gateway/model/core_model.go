package model

type ResponseGetCore struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
