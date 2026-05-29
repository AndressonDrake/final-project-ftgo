package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type branchRepository struct {
	db *gorm.DB
}

func BranchRepository(db *gorm.DB) domain.BranchRepository {
	return &branchRepository{db: db}
}

func (b *branchRepository) Begin() (tx *gorm.DB) {
	tx = b.db.Begin()
	return
}

func (b *branchRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (b *branchRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (b *branchRepository) Create(tx *gorm.DB, request model.Branch) (err error) {
	err = tx.Create(&request).Error
	return
}