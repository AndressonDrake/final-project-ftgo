package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type branchUsecase struct {
	branchRepository domain.BranchRepository
}

func BranchUsecase(branchRepository domain.BranchRepository) domain.BranchUsecase {
	return &branchUsecase{branchRepository: branchRepository}
}

func (b *branchUsecase) Create(request model.CreateBranch) (err error) {
	var req model.Branch
	req.NamaCabang = request.NamaCabang
	req.Alamat = request.Alamat

	tx := b.branchRepository.Begin()
	err = b.branchRepository.Create(tx, req)
	if err != nil {
		b.branchRepository.Rollback(tx)
		return
	}
	b.branchRepository.Commit(tx)
	return
}