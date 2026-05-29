package domain

import (
	"github.com/labstack/echo/v4"
	"gw-heatlcare.com/model"
)

type PostUsecase interface {
	Post(request model.RequestGeneral, email string) (message, detail string, err error)
}

type PostHandler interface {
	Post(c echo.Context) (err error)
}
