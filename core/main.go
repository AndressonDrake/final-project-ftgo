package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"core-healtcare.com/config"
	"core-healtcare.com/domain"
	l "core-healtcare.com/helper/logger"
	"core-healtcare.com/model"
	"core-healtcare.com/repository"
	"core-healtcare.com/usecase"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	MedicineUsecase domain.MedicineUsecase
)

// @title Core
// @version 1.0
// @description API Documentation P2
// @host localhost:12100
// @BasePath /core/

// @securityDefinitions.apikey api-key
// @in header
// @name Authorization
func main() {
	l.NewLogger("core-health")

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
		panic(err)
	}

	DB_NAME := os.Getenv("DB_NAME")
	DB_USER := os.Getenv("DB_USER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	DB_PASSWORD := os.Getenv("DB_PASSWORD")

	REDIS_HOST := os.Getenv("REDIS_HOST")
	REDIS_PORT, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))

	// API_KEY := os.Getenv("API_KEY")

	url := fmt.Sprintf("0.0.0.0:%d", PORT)

	db, err := config.ConnectDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	rdb, err := config.ConnectRedis(REDIS_HOST, REDIS_PASSWORD, REDIS_PORT)

	medicineRepository := repository.MedicineRepository(db)

	MedicineUsecase = usecase.MedicineUsecase(medicineRepository)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		RedisConsumer(ctx, rdb, "HEALTHCARE:STREAM", "HEALTHCARE:GROUP", "HEALTHCARE:CONSUMER")
	}()

	e.Start(url)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		e.Shutdown(ctx)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	cancel()
	wg.Wait()

}

func RedisConsumer(ctx context.Context, rdb *redis.Client, stream, groupName, consumerName string) {
	var reqRedis model.Consumer

	err := rdb.FlushAll(ctx).Err()
	if err != nil {
		l.Log.Error(l.Fields{
			"error": err.Error(),
		}, nil, err.Error())
		return
	}

	groups, err := rdb.XInfoGroups(ctx, stream).Result()
	if err != nil && err != redis.Nil {
		l.Log.Error(l.Fields{
			"error": err.Error(),
		}, nil, err.Error())
		
	}

	groupExists := false
	for _, group := range groups {
		if group.Name == groupName {
			groupExists = true
			break
		}
	}

	if !groupExists {
		_, err = rdb.XGroupCreateMkStream(ctx, stream, groupName, "0").Result()
		if err != nil && err != redis.Nil {
			if err.Error() != "BUSYGROUP Consumer Group name already exists" {
				l.Log.Error(l.Fields{
					"error": err.Error(),
				}, nil, err.Error())
				return
			}

		}
	}

	log.Println("Ara ara listening")

	streams := []string{stream, ">"}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			xs, err := rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    groupName,
				Consumer: consumerName,
				Streams:  streams,
				Count:    1,
				Block:    0, // Block indefinitely until a message arrives
			}).Result()

			if err != nil {
				if err == context.Canceled {
					return
				}
				// fmt.Printf("Error reading from stream: %v\n", err)

				time.Sleep(1 * time.Second)
				continue
			}

			for _, _stream := range xs {
				for _, message := range _stream.Messages {
					fmt.Printf("Received message: %v\n", message.Values["data"])

					messageValue, ok := message.Values["data"].(string)

					if !ok {
						l.Log.Error(l.Fields{
							"error": "error parse",
						}, nil, "error parse")
						continue
					}

					err := json.Unmarshal([]byte(messageValue), &reqRedis)

					if err != nil {
						continue
					}

					if reqRedis.TrxType == "MEDICINE" {
						if reqRedis.SubType == "CREATE" {
							var req model.CreateMedicine

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {

								continue
							}

							l.Log.Info(l.Fields{
								"data": req,
							}, nil, "info data create")

							err = MedicineUsecase.Create(req)
							if err != nil {
								l.Log.Error(l.Fields{
									"error":err.Error(),
								},nil,"error create medicine")
								continue
							}
						}
					}else if reqRedis.TrxType == "APOINTMENT"{
						if reqRedis.SubType == "CREATE"{
							//panggil usecase
						}
					}else if reqRedis.TrxType == "MEDICAL-RECORD"{
						if reqRedis.SubType == "CREATE"{

						}
					}


					rdb.XAck(ctx, stream, groupName, message.ID)

				}
			}
		}

	}

}
