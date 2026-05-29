package model

type CreateUser struct {
	IdRole   int    `json:"id_role"`
	IdCabang int    `json:"id_cabang"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	NoHp     string `json:"no_hp"`
}