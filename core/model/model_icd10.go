package model

type ResponseSuccessGetICD10 struct {
	Message string  `json:"message"`
	Data    []ICD10 `json:"data"`
}
