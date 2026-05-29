package repository

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Begin() (tx *gorm.DB) {
	tx = u.db.Begin()
	return
}

func (u *userRepository) Commit(tx *gorm.DB) (err error) {
	err = tx.Commit().Error
	return
}

func (u *userRepository) Rollback(tx *gorm.DB) (err error) {
	err = tx.Rollback().Error
	return
}

func (u *userRepository) Create(tx *gorm.DB, request model.User) (err error) {
	err = tx.Create(&request).Error
	return
}

func (u *userRepository) FindAll() (response []model.User, err error) {
	err = u.db.Find(&response).Error
	return
}

func (u *userRepository) FindById(id int) (response model.User, err error) {
	err = u.db.Where("id_user = ?", id).First(&response).Error
	return
}

func (u *userRepository) Update(tx *gorm.DB, request model.User) (err error) {
	err = tx.Save(&request).Error
	return
}

func (u *userRepository) Delete(tx *gorm.DB, id int) (err error) {
	err = tx.Where("id_user = ?", id).Delete(&model.User{}).Error
	return
}