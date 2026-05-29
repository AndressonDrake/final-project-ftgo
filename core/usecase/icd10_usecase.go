package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type icd10Usecase struct {
	icd10Repository domain.ICD10Repository
}

func ICD10Usecase(icd10Repository domain.ICD10Repository) domain.ICD10Usecase {
	return &icd10Usecase{icd10Repository: icd10Repository}
}

func (i *icd10Usecase) Create(request model.CreateICD10) (err error) {

	req := model.ICD10{
		KodeICD:      request.KodeIcd,
		NamaPenyakit: request.NamaPenyakit,
	}

	tx := i.icd10Repository.Begin()

	err = i.icd10Repository.Create(tx, req)
	if err != nil {
		i.icd10Repository.Rollback(tx)
		return
	}

	i.icd10Repository.Commit(tx)

	return
}

func (u *icd10Usecase) Get() (data []model.ICD10, message, detail string, err error) {
	data, err = u.icd10Repository.Get()
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get icd10"
	return
}

func (u *icd10Usecase) GetByID(id int) (data model.ICD10, message, detail string, err error) {
	data, err = u.icd10Repository.FindByID(id)
	if err != nil {
		detail = err.Error()
		message = "data not found"
		return
	}
	message = "success get icd10 by ID"
	return
}