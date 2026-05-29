package model

type User struct {
	IDUser    int    `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"`
	IDRole    int    `gorm:"column:id_role" json:"id_role"`
	IDCabang  int    `gorm:"column:id_cabang" json:"id_cabang"`
	Nama      string `gorm:"column:nama;type:varchar(150)" json:"nama"`
	Email     string `gorm:"column:email;type:varchar(150);unique" json:"email"`
	Password  string `gorm:"column:password;type:varchar(255)" json:"password"`
	NoHP      string `gorm:"column:no_hp;type:varchar(20)" json:"no_hp"`
}

func (User) TableName() string {
	return "users"
}