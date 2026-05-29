package domain

import (
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MedicineRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)
	Create(tx *gorm.DB, request model.Medicine) (err error)
	Get() (data []model.Medicine, err error)
	FindByID(id int) (data model.Medicine, err error) // 🆕
}

type MedicineUsecase interface {
	Create(request model.CreateMedicine) (err error)
	Get() (data []model.Medicine, message, detail string, err error)
	GetByID(id int) (data model.Medicine, message, detail string, err error) // 🆕
}

type MedicineHandler interface {
	Get(c echo.Context) (err error)
	GetByID(c echo.Context) (err error) // 🆕
}
