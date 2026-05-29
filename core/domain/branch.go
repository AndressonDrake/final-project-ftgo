package domain

import (
	"core-healtcare.com/model"
	"gorm.io/gorm"
)

type BranchRepository interface {
	Begin() (tx *gorm.DB)
	Commit(tx *gorm.DB) (err error)
	Rollback(tx *gorm.DB) (err error)

	Create(tx *gorm.DB, request model.Branch) (err error)
}

type BranchUsecase interface {
	Create(request model.CreateBranch) (err error)
}

type BranchHandler interface{
	
}

