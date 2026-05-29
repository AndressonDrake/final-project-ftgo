package model

type CreateRole struct {
	IdRole   int    `json:"id_role" gorm:"primaryKey"`
	NamaRole string `json:"nama_role"`
}
