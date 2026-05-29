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