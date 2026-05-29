package model

type CreateHealthNews struct {
	Judul          string `json:"judul"`
	Sumber         string `json:"sumber"`
	Kategori       string `json:"kategori"`
	TanggalPublish string `json:"tanggal_publish"`
	Url            string `json:"url"`
}

type ResponseSuccessGetHealthNews struct {
	Message string       `json:"message"`
	Data    []HealthNews `json:"data"`
}

type ResponseSuccessGetHealthNewsByID struct {
	Message string     `json:"message"`
	Data    HealthNews `json:"data"`
}