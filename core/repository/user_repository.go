package repository

import "gorm.io/gorm"

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Begin()(tx *gorm.DB){
	tx = u.db.Begin()
	return
}

func (u *userRepository) Commit(tx *gorm.DB)(err error){
	err = tx.Commit().Error
	return
}

func (u *userRepository) Rollback(tx *gorm.DB)(err error){
	err = tx.Rollback().Error
	return
}
