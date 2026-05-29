package model

import (
	"encoding/json"
	"time"
)

type Role struct {
	IDRole   int    `gorm:"column:id_role;primaryKey;autoIncrement" json:"id_role"`
	NamaRole string `gorm:"column:nama_role;type:varchar(100)" json:"nama_role"`
}

func (Role) TableName() string {
	return "roles"
}

type Branch struct {
	IDCabang   int    `gorm:"column:id_cabang;primaryKey;autoIncrement" json:"id_cabang"`
	NamaCabang string `gorm:"column:nama_cabang;type:varchar(150)" json:"nama_cabang"`
	Alamat     string `gorm:"column:alamat;type:text" json:"alamat"`
}

func (Branch) TableName() string {
	return "branches"
}

type User struct {
	IDUser   int    `gorm:"column:id_user;primaryKey;autoIncrement" json:"id_user"`
	IDRole   int    `gorm:"column:id_role" json:"id_role"`
	IDCabang int    `gorm:"column:id_cabang" json:"id_cabang"`
	Nama     string `gorm:"column:nama;type:varchar(150)" json:"nama"`
	Email    string `gorm:"column:email;type:varchar(150);unique" json:"email"`
	Password string `gorm:"column:password;type:varchar(255)" json:"password"`
	NoHP     string `gorm:"column:no_hp;type:varchar(20)" json:"no_hp"`

	Role   Role   `gorm:"foreignKey:IDRole;references:IDRole" json:"role"`
	Branch Branch `gorm:"foreignKey:IDCabang;references:IDCabang" json:"branch"`
}

func (User) TableName() string {
	return "users"
}

type PatientStatus struct {
	IDStatus   int    `gorm:"column:id_status;primaryKey;autoIncrement" json:"id_status"`
	NamaStatus string `gorm:"column:nama_status;type:varchar(100)" json:"nama_status"`
}

func (PatientStatus) TableName() string {
	return "patient_status"
}

type Patient struct {
	IDPatient     int       `gorm:"column:id_patient;primaryKey;autoIncrement" json:"id_patient"`
	IDStatus      int       `gorm:"column:id_status" json:"id_status"`
	Nama          string    `gorm:"column:nama;type:varchar(150)" json:"nama"`
	NIK           string    `gorm:"column:nik;type:varchar(30);unique" json:"nik"`
	TanggalLahir  time.Time `gorm:"column:tanggal_lahir;type:date" json:"tanggal_lahir"`
	Gender        string    `gorm:"column:gender;type:varchar(20)" json:"gender"`
	Alamat        string    `gorm:"column:alamat;type:text" json:"alamat"`
	NoHP          string    `gorm:"column:no_hp;type:varchar(20)" json:"no_hp"`
	GolonganDarah string    `gorm:"column:golongan_darah;type:varchar(5)" json:"golongan_darah"`

	Status PatientStatus `gorm:"foreignKey:IDStatus;references:IDStatus" json:"status"`
}

func (Patient) TableName() string {
	return "patients"
}

type Medicine struct {
	IDObat      int       `gorm:"column:id_obat;primaryKey;autoIncrement" json:"id_obat"`
	NamaObat    string    `gorm:"column:nama_obat;type:varchar(150)" json:"nama_obat"`
	Kategori    string    `gorm:"column:kategori;type:varchar(100)" json:"kategori"`
	Stok        int       `gorm:"column:stok" json:"stok"`
	Harga       float64   `gorm:"column:harga;type:decimal(12,2)" json:"harga"`
	ExpiredDate time.Time `gorm:"column:expired_date;type:date" json:"expired_date"`
}

func (Medicine) TableName() string {
	return "medicines"
}

type ICD10 struct {
	IDICD        int    `gorm:"column:id_icd;primaryKey;autoIncrement" json:"id_icd"`
	KodeICD      string `gorm:"column:kode_icd;type:varchar(20)" json:"kode_icd"`
	NamaPenyakit string `gorm:"column:nama_penyakit;type:varchar(200)" json:"nama_penyakit"`
	Kategori     string `gorm:"column:kategori;type:varchar(150)" json:"kategori"`
}

func (ICD10) TableName() string {
	return "icd10"
}

type Appointment struct {
	IDAppointment int       `gorm:"column:id_appointment;primaryKey;autoIncrement" json:"id_appointment"`
	IDPatient     int       `gorm:"column:id_patient" json:"id_patient"`
	IDDoctor      int       `gorm:"column:id_doctor" json:"id_doctor"`
	Tanggal       time.Time `gorm:"column:tanggal;type:timestamp" json:"tanggal"`
	Keluhan       string    `gorm:"column:keluhan;type:text" json:"keluhan"`
	TekananDarah  string    `gorm:"column:tekanan_darah;type:varchar(20)" json:"tekanan_darah"`
	SuhuTubuh     float64   `gorm:"column:suhu_tubuh;type:decimal(4,1)" json:"suhu_tubuh"`
	BeratBadan    float64   `gorm:"column:berat_badan;type:decimal(5,2)" json:"berat_badan"`
	Status        string    `gorm:"column:status;type:varchar(50)" json:"status"`

	Patient Patient `gorm:"foreignKey:IDPatient;references:IDPatient" json:"patient"`
	Doctor  User    `gorm:"foreignKey:IDDoctor;references:IDUser" json:"doctor"`
}

func (Appointment) TableName() string {
	return "appointments"
}

type MedicalRecord struct {
	IDRecord       int    `gorm:"column:id_record;primaryKey;autoIncrement" json:"id_record"`
	IDAppointment  int    `gorm:"column:id_appointment" json:"id_appointment"`
	IDICD          int    `gorm:"column:id_icd" json:"id_icd"`
	HasilLab       string `gorm:"column:hasil_lab;type:text" json:"hasil_lab"`
	HasilRadiologi string `gorm:"column:hasil_radiologi;type:text" json:"hasil_radiologi"`
	Tindakan       string `gorm:"column:tindakan;type:text" json:"tindakan"`
	Catatan        string `gorm:"column:catatan;type:text" json:"catatan"`

	Appointment Appointment `gorm:"foreignKey:IDAppointment;references:IDAppointment" json:"appointment"`
	ICD10       ICD10       `gorm:"foreignKey:IDICD;references:IDICD" json:"icd10"`
}

func (MedicalRecord) TableName() string {
	return "medical_records"
}

type Prescription struct {
	IDResep     int    `gorm:"column:id_resep;primaryKey;autoIncrement" json:"id_resep"`
	IDRecord    int    `gorm:"column:id_record" json:"id_record"`
	IDObat      int    `gorm:"column:id_obat" json:"id_obat"`
	Jumlah      int    `gorm:"column:jumlah" json:"jumlah"`
	AturanPakai string `gorm:"column:aturan_pakai;type:text" json:"aturan_pakai"`

	MedicalRecord MedicalRecord `gorm:"foreignKey:IDRecord;references:IDRecord" json:"medical_record"`
	Medicine      Medicine      `gorm:"foreignKey:IDObat;references:IDObat" json:"medicine"`
}

func (Prescription) TableName() string {
	return "prescriptions"
}

type Payment struct {
	IDPayment        int     `gorm:"column:id_payment;primaryKey;autoIncrement" json:"id_payment"`
	IDAppointment    int     `gorm:"column:id_appointment" json:"id_appointment"`
	Total            float64 `gorm:"column:total;type:decimal(12,2)" json:"total"`
	MetodePembayaran string  `gorm:"column:metode_pembayaran;type:varchar(50)" json:"metode_pembayaran"`
	StatusPembayaran string  `gorm:"column:status_pembayaran;type:varchar(50)" json:"status_pembayaran"`

	Appointment Appointment `gorm:"foreignKey:IDAppointment;references:IDAppointment" json:"appointment"`
}

func (Payment) TableName() string {
	return "payments"
}

type DiseaseMonitoring struct {
	IDMonitoring  int       `gorm:"column:id_monitoring;primaryKey;autoIncrement" json:"id_monitoring"`
	IDICD         int       `gorm:"column:id_icd" json:"id_icd"`
	Negara        string    `gorm:"column:negara;type:varchar(100)" json:"negara"`
	TotalKasus    int       `gorm:"column:total_kasus" json:"total_kasus"`
	TotalKematian int       `gorm:"column:total_kematian" json:"total_kematian"`
	TotalSembuh   int       `gorm:"column:total_sembuh" json:"total_sembuh"`
	TanggalUpdate time.Time `gorm:"column:tanggal_update;type:timestamp" json:"tanggal_update"`

	ICD10 ICD10 `gorm:"foreignKey:IDICD;references:IDICD" json:"icd10"`
}

func (DiseaseMonitoring) TableName() string {
	return "disease_monitoring"
}

type HealthNews struct {
	IDNews         int       `gorm:"column:id_news;primaryKey;autoIncrement" json:"id_news"`
	Judul          string    `gorm:"column:judul;type:varchar(255)" json:"judul"`
	Sumber         string    `gorm:"column:sumber;type:varchar(150)" json:"sumber"`
	Kategori       string    `gorm:"column:kategori;type:varchar(100)" json:"kategori"`
	TanggalPublish time.Time `gorm:"column:tanggal_publish;type:timestamp" json:"tanggal_publish"`
	URL            string    `gorm:"column:url;type:text" json:"url"`
}

func (HealthNews) TableName() string {
	return "health_news"
}

type Consumer struct {
	TrxType string          `json:"trx_type"`
	SubType string          `json:"sub_type"`
	Data    json.RawMessage `json:"data"`
	Email   string          `json:"email"`
}
