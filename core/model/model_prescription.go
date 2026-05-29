package model

type CreatePrescription struct {
	IdRecord    int    `json:"id_record"`
	IdObat      int    `json:"id_obat"`
	Jumlah      int    `json:"jumlah"`
	AturanPakai string `json:"aturan_pakai"`
}