package model

type CreateMedicine struct {
	NamaObat    string  `json:"nama_obat"`
	Kategori    string  `json:"kategori"`
	Stok        int     `json:"stok"`
	Harga       float64 `json:"harga"`
	ExpiredDate string  `json:"expired_date"`
}

type ResponseSuccessGetMedicine struct {
	Message string     `json:"message"`
	Data    []Medicine `json:"data"`
}

type ResponseError struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}
