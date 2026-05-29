package model


type CreatePatient struct {
	IdStatus      int    `json:"id_status"`
	Nama          string `json:"nama"`
	Nik           string `json:"nik"`
	TanggalLahir  string `json:"tanggal_lahir"`
	Gender        string `json:"gender"`
	Alamat        string `json:"alamat"`
	NoHp          string `json:"no_hp"`
	GolonganDarah string `json:"golongan_darah"`
}

type ResponseSuccessGetPatient struct {
	Message string    `json:"message"`
	Data    []Patient `json:"data"`
}

type ResponseSuccessGetPatientByID struct {
	Message string  `json:"message"`
	Data    Patient `json:"data"`
}