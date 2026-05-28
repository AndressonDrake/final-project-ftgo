package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type medicineRepository struct {
	db *gorm.DB
}

func MedicineRepository(db *gorm.DB) domain.MedicineRepository {
	return &medicineRepository{db: db}
}

func (m *medicineRepository) Begin() (tx *gorm.DB) {
	tx = m.db.Begin()
	return
}

func (m *medicineRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (m *medicineRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (m *medicineRepository) Create(tx *gorm.DB, request model.Medicine) (err error) {
	err = tx.Create(&request).Error
	return
}


func (m *medicineRepository)Get()(data []model.Medicine,err error){
	err = m.db.Find(&data).Error
	return
}