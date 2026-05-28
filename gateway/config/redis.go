package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(host, password string, port int) (rdb *redis.Client, err error) {

	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
		Username: "default",
	})

	ctx := context.Background()


	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		remark := fmt.Sprintf("error connect to redis %[1]v", err)
		log.Println(remark)
		return
	}

	return
}
