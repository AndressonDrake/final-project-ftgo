package domain

import (
	"github.com/labstack/echo/v4"
	"gw-heatlcare.com/model"
)

type UserRepository interface {
	Login(email string) (data model.User, err error)
}

type UserUsecase interface {
	Login(request model.ReqeustLoginUser) (token, message, detail string, err error)
}

type UserHandler interface{
	Login(c echo.Context) (err error)
}