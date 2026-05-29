package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type HealthNewsRepository interface {
	Get() (data []model.HealthNews, err error)
	FindByID(id int) (data model.HealthNews, err error)
}
type HealthNewsUsecase interface {
	Get() (data []model.HealthNews, message, detail string, err error)
	GetByID(id int) (data model.HealthNews, message, detail string, err error)
}
type HealthNewsHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error)
}
