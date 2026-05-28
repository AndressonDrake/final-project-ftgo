package model

type ResponseSuccessGetHealthNews struct {
	Message string       `json:"message"`
	Data    []HealthNews `json:"data"`
}
