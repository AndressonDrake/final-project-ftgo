package usecase

import (
	"gw-heatlcare.com/domain"
	l "gw-heatlcare.com/helper/logger"
	"gw-heatlcare.com/model"
)

type postUsecase struct {
	redisRepository domain.RedisRepository
}

func PostUsecase(redisRepository domain.RedisRepository) domain.PostUsecase {
	return &postUsecase{redisRepository: redisRepository}
}

func (p *postUsecase) Post(request model.RequestGeneral,email string) (message, detail string, err error) {
	var reqProducer model.RedisProducer

	reqProducer.TrxType = request.TrxType
	reqProducer.SubType = request.SubType
	reqProducer.Data = request.Data
	reqProducer.Email = email

	err = p.redisRepository.RedisProducer(reqProducer)
	if err != nil {
		detail = err.Error()
		message = "internal server error"
		l.Log.Error(l.Fields{
			"error": err.Error(),
		}, nil, "error produce message")
		return
	}

	message = "success post data"
	return
}
