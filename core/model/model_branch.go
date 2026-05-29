package model

type CreateBranch struct {
	IdCabang   int    `json:"id_cabang" gorm:"primaryKey"`
	NamaCabang string `json:"nama_cabang"`
	Alamat     string `json:"alamat"`
}

