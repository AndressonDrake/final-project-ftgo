package domain

import "github.com/labstack/echo/v4"

type CoreRepository interface {
	GetCore(pathUrl string) (response interface{}, err error)
}

type CoreUsecase interface {
	GetCore(pathUrl string) (data interface{}, message, detail string, err error)
}

type CoreHandler interface {
	GetCore(e echo.Context) (err error)
}
