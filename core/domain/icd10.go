package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ICD10Repository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.ICD10) (err error)
	Get() (data []model.ICD10, err error)
	FindByID(id int) (data model.ICD10, err error)
}

type ICD10Usecase interface {
	Create(request model.CreateICD10) (err error)
	Get() (data []model.ICD10, message, detail string, err error)
	GetByID(id int) (data model.ICD10, message, detail string, err error)
}

type ICD10Handler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}