package model

type CreateHealthNews struct {
	Judul          string `json:"judul"`
	Sumber         string `json:"sumber"`
	Kategori       string `json:"kategori"`
	TanggalPublish string `json:"tanggal_publish"`
	Url            string `json:"url"`
}