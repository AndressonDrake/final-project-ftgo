package model

type ResponseSuccessGetICD10 struct {
	Message string  `json:"message"`
	Data    []ICD10 `json:"data"`
}

type ResponseSuccessGetICD10ByID struct {
	Message string `json:"message"`
	Data    ICD10  `json:"data"`
}
