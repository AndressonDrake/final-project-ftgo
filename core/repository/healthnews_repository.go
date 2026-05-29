package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type healthNewsRepository struct {
	db *gorm.DB
}

func HealthNewsRepository(db *gorm.DB) domain.HealthNewsRepository {
	return &healthNewsRepository{db: db}
}

func (h *healthNewsRepository) Begin() (tx *gorm.DB) {
	tx = h.db.Begin()
	return
}

func (h *healthNewsRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (h *healthNewsRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (h *healthNewsRepository) Create(tx *gorm.DB, request model.HealthNews) (err error) {
	err = tx.Create(&request).Error
	return
}
func (r *healthNewsRepository) Get() (data []model.HealthNews, err error) {
	err = r.db.Find(&data).Error
	return
}

func (r *healthNewsRepository) FindByID(id int) (data model.HealthNews, err error) {
	err = r.db.First(&data, id).Error
	return
}