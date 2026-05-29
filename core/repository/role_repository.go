package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func RoleRepository(db *gorm.DB) domain.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) Begin() (tx *gorm.DB) {
	tx = r.db.Begin()
	return
}

func (r *roleRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (r *roleRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (r *roleRepository) Create(tx *gorm.DB, request model.Role) (err error) {
	err = tx.Create(&request).Error
	return
}