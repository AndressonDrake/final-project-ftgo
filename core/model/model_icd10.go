package model


type CreateICD10 struct {
	KodeIcd      string `json:"kode_icd"`
	NamaPenyakit string `json:"nama_penyakit"`
}

type ResponseSuccessGetICD10 struct {
	Message string  `json:"message"`
	Data    []ICD10 `json:"data"`
}

type ResponseSuccessGetICD10ByID struct {
	Message string `json:"message"`
	Data    ICD10  `json:"data"`
}