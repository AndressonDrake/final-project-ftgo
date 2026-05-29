package usecase

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
)

type roleUsecase struct {
	roleRepository domain.RoleRepository
}

func RoleUsecase(roleRepository domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{roleRepository: roleRepository}
}

func (r *roleUsecase) Create(request model.CreateRole) (err error) {

	req := model.Role{
		NamaRole: request.NamaRole,
	}

	tx := r.roleRepository.Begin()

	err = r.roleRepository.Create(tx, req)
	if err != nil {
		r.roleRepository.Rollback(tx)
		return
	}

	r.roleRepository.Commit(tx)

	return
}