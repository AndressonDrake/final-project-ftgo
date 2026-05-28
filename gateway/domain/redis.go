package domain

import "gw-heatlcare.com/model"

type RedisRepository interface {
	RedisProducer(message model.RedisProducer) (err error)
}