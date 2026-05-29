package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gw-heatlcare.com/config"
	_ "gw-heatlcare.com/docs"
	"gw-heatlcare.com/handler"
	"gw-heatlcare.com/helper"
	l "gw-heatlcare.com/helper/logger"
	"gw-heatlcare.com/repository"
	"gw-heatlcare.com/usecase"
)

func main() {
	l.NewLogger("gw-health")

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

	helper.SECRETKEY = os.Getenv("SECRET")

	API_KEY := os.Getenv("API_KEY")

	CORE_URL := os.Getenv("CORE_URL")

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))

	url := fmt.Sprintf("0.0.0.0:%d", PORT)

	rdb, err := config.ConnectRedis(REDIS_HOST, REDIS_PASSWORD, REDIS_PORT)

	db, err := config.ConnectDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	redisRepository := repository.RedisRepository(rdb, "HEALTHCARE:STREAM")

	coreRepository := repository.CoreRepository(API_KEY, CORE_URL)

	userRepository := repository.UserRepository(db)

	postUsecase := usecase.PostUsecase(redisRepository)

	coreUsecase := usecase.GetUsecase(coreRepository)

	userUsecase := usecase.UserUsecase(userRepository)

	postHandler := handler.PostHandler(postUsecase)

	coreHandler := handler.CoreHandler(coreUsecase)

	userHandler := handler.UserHandler(userUsecase)

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.POST("/api", postHandler.Post, handler.Middleware)

	e.GET("/api", coreHandler.GetCore, handler.Middleware)

	e.POST("/login", userHandler.Login)
	e.POST("/register", postHandler.PostCreateUser)

	e.Start(url)

}
