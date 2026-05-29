package usecase

import "gw-heatlcare.com/domain"

type coreUsecase struct {
	coreRepository domain.CoreRepository
}

func GetUsecase(coreRepository domain.CoreRepository) domain.CoreUsecase {
	return &coreUsecase{coreRepository: coreRepository}
}

func (c *coreUsecase) GetCore(pathUrl string) (data interface{}, message, detail string, err error) {
	data, err = c.coreRepository.GetCore(pathUrl)
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		return
	}
	message = "success get core"
	return
}
