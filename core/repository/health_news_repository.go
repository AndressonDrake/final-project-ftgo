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

func (r *healthNewsRepository) Get() (data []model.HealthNews, err error) {
	err = r.db.Find(&data).Error
	return
}
