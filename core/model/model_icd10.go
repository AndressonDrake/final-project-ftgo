package model


type CreateICD10 struct {
	KodeIcd      string `json:"kode_icd"`
	NamaPenyakit string `json:"nama_penyakit"`
}
