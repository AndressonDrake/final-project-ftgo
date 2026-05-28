package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type ICD10Repository interface {
	Get() (data []model.ICD10, err error)
}

type ICD10Usecase interface {
	Get() (data []model.ICD10, message, detail string, err error)
}

type ICD10Handler interface {
	Get(c echo.Context) (err error)
}
