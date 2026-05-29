package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HealthNewsRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.HealthNews) (err error)
	Get() (data []model.HealthNews, err error)
	FindByID(id int) (data model.HealthNews, err error)
}

type HealthNewsUsecase interface {
	Create(request model.CreateHealthNews) (err error)
	Get() (data []model.HealthNews, message, detail string, err error)
	GetByID(id int) (data model.HealthNews, message, detail string, err error)
	
}

type HealthNewsHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}