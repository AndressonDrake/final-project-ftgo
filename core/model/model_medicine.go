package model



type CreateMedicine struct {
	NamaObat    string  `json:"nama_obat"`
	Kategori    string  `json:"kategori"`
	Stok        int     `json:"stok"`
	Harga       float64 `json:"harga"`
	ExpiredDate string  `json:"expired_date"`
}
