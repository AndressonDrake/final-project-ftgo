package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/helper"
	"core-healtcare.com/model"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func UserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Create(request model.CreateUser) (err error) {

	req := model.User{
		IDRole:   request.IdRole,
		IDCabang: request.IdCabang,
		Nama:     request.Nama,
		Email:    request.Email,
		Password: helper.HashPassword(request.Password),
		NoHP:     request.NoHp,
	}

	tx := u.userRepository.Begin()

	err = u.userRepository.Create(tx, req)
	if err != nil {
		u.userRepository.Rollback(tx)
		return
	}

	u.userRepository.Commit(tx)

	return
}
