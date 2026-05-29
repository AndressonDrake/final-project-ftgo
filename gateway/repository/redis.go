package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"gw-heatlcare.com/domain"
	l "gw-heatlcare.com/helper/logger"
	"gw-heatlcare.com/model"
)

type redisRepository struct {
	rdb    *redis.Client
	stream string
}

func RedisRepository(rdb *redis.Client, stream string) domain.RedisRepository {
	return &redisRepository{rdb: rdb, stream: stream}
}

func (r *redisRepository) RedisProducer(message model.RedisProducer) (err error) {
	ctx := context.Background()
	msgJson, err := json.Marshal(message)
	if err != nil {
		remark := "error marshall message"
		l.Log.Error(l.Fields{
			"error": err.Error(),
		}, message, remark)
		return
	}

	err = r.rdb.XAdd(ctx, &redis.XAddArgs{
		Stream: r.stream,
		Values: map[string]interface{}{"data": string(msgJson)},
	}).Err()
	if err != nil {
		remark := "error send redis stream"
		l.Log.Error(l.Fields{
			"error": err.Error(),
		}, message, remark)
		return
	}

	return
}

