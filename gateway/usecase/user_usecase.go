package usecase

import (
	"fmt"

	"gw-heatlcare.com/domain"
	"gw-heatlcare.com/helper"
	"gw-heatlcare.com/model"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func UserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) Login(request model.ReqeustLoginUser) (token, message, detail string, err error) {

	email := request.Email
	password := request.Password

	dataUser, err := u.userRepository.Login(email)
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}

	isverified := helper.VerifyPassword(password, dataUser.Password)

	if !isverified {
		err = fmt.Errorf("error login")
		message = "internal server error"
		return
	}

	token, _ = helper.GenerateToken(dataUser.IDUser, dataUser.Email)

	message = "success login"

	return
}
