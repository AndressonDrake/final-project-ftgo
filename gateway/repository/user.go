package repository

import (
	"gorm.io/gorm"
	"gw-heatlcare.com/domain"
	"gw-heatlcare.com/model"
)

type userRepository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Login(email string) (data model.User, err error) {
	err = u.db.Where("email = ? ", email).Find(&data).Error
	return
}
