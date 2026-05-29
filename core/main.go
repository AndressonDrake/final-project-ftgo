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
	"core-healtcare.com/handler"
	"core-healtcare.com/helper"
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
	MedicineUsecase          domain.MedicineUsecase
	AppointmentUsecase       domain.AppointmentUsecase
	BranchUsecase            domain.BranchUsecase
	DiseaseMonitoringUsecase domain.DiseaseMonitoringUsecase
	HealthNewsUsecase        domain.HealthNewsUsecase
	MedicalRecordUsecase     domain.MedicalRecordUsecase
	PatientUsecase           domain.PatientUsecase
	UserUsecase              domain.UserUsecase
	ICD10Usecase             domain.ICD10Usecase
	PrescriptionUsecase      domain.PrescriptionUsecase
	PaymentUsecase           domain.PaymentUsecase
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

	helper.URLMAILING = os.Getenv("URL_MAILING")
	helper.XAPIKey = os.Getenv("XAPIKEY")

	PORT, _ := strconv.Atoi(os.Getenv("PORT"))

	API_KEY := os.Getenv("API_KEY")

	handler.API_KEY = API_KEY

	url := fmt.Sprintf("0.0.0.0:%d", PORT)

	db, err := config.ConnectDB(DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	rdb, err := config.ConnectRedis(REDIS_HOST, REDIS_PASSWORD, REDIS_PORT)

	medicineRepository := repository.MedicineRepository(db)

	MedicineUsecase = usecase.MedicineUsecase(medicineRepository)

	medicineHandler := handler.MedicineHandler(MedicineUsecase)

	appointmentRepository := repository.AppointmentRepository(db)
	AppointmentUsecase = usecase.AppointmentUsecase(appointmentRepository)
	appointmentHandler := handler.AppointmentHandler(AppointmentUsecase)

	branchRepository := repository.BranchRepository(db)
	BranchUsecase = usecase.BranchUsecase(branchRepository)

	diseaseRepository := repository.DiseaseMonitoringRepository(db)
	DiseaseMonitoringUsecase = usecase.DiseaseMonitoringUsecase(diseaseRepository)
	diseaseMonitoringHandler := handler.DiseaseMonitoringHandler(DiseaseMonitoringUsecase)

	newsRepository := repository.HealthNewsRepository(db)
	HealthNewsUsecase = usecase.HealthNewsUsecase(newsRepository)
	healthNewsHandler := handler.HealthNewsHandler(HealthNewsUsecase)

	icdRepository := repository.ICD10Repository(db)
	ICD10Usecase = usecase.ICD10Usecase(icdRepository)
	icd10Handler := handler.ICD10Handler(ICD10Usecase)

	prescriptionRepository := repository.PrescriptionRepository(db)
	PrescriptionUsecase = usecase.PrescriptionUsecase(prescriptionRepository)
	prescriptionHandler := handler.PrescriptionHandler(PrescriptionUsecase)

	paymentRepository := repository.PaymentRepository(db)
	PaymentUsecase = usecase.PaymentUsecase(paymentRepository)
	paymentHandler := handler.PaymentHandler(PaymentUsecase)

	medicalRecordRepository := repository.MedicalRecordRepository(db)
	MedicalRecordUsecase = usecase.MedicalRecordUsecase(medicalRecordRepository)
	medicalRecordHandler := handler.MedicalRecordHandler(MedicalRecordUsecase)

	patientRepository := repository.PatientRepository(db)
	PatientUsecase = usecase.PatientUsecase(patientRepository)
	patientHandler := handler.PatientHandler(PatientUsecase)

	userRepository := repository.UserRepository(db)
	UserUsecase = usecase.UserUsecase(userRepository)
	


	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/api/medicine", medicineHandler.Get, handler.Middleware)

	e.GET("/api/appointment", appointmentHandler.Get, handler.Middleware)

	e.GET("/api/disease-monitoring", diseaseMonitoringHandler.Get, handler.Middleware)

	e.GET("/api/health-news", healthNewsHandler.Get, handler.Middleware)

	e.GET("/api/icd10", icd10Handler.Get, handler.Middleware)

	e.GET("/api/prescription", prescriptionHandler.Get, handler.Middleware)

	e.GET("/api/payment", paymentHandler.Get, handler.Middleware)

	e.GET("/api/patient", patientHandler.Get, handler.Middleware)

	e.GET("/api/medical-record", medicalRecordHandler.Get, handler.Middleware)

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

							err = MedicineUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "APPOINTMENT" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateAppointment

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								fmt.Printf("Received message: %v\n", message.Values["data"])
								continue
							}
							fmt.Printf("AppointmentUsecase: %#v\n", AppointmentUsecase)

							err = AppointmentUsecase.Create(req)
							if err != nil {
								fmt.Println("APPOINTMENT ERROR:", err)
								continue
							}
						}

					} else if reqRedis.TrxType == "MEDICAL_RECORD" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateMedicalRecord

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = MedicalRecordUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "PATIENT" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreatePatient

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = PatientUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "USER" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateUser

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = UserUsecase.Create(req)
							if err != nil {
								continue
							}
						}
					} else if reqRedis.TrxType == "ICD10" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateICD10

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = ICD10Usecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "PRESCRIPTION" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreatePrescription

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = PrescriptionUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "PAYMENT" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreatePayment

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = PaymentUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "DISEASE_MONITORING" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateDiseaseMonitoring

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = DiseaseMonitoringUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "HEALTH_NEWS" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateHealthNews

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = HealthNewsUsecase.Create(req)
							if err != nil {
								continue
							}
						}

					} else if reqRedis.TrxType == "BRANCH" {

						if reqRedis.SubType == "CREATE" {

							var req model.CreateBranch

							err = json.Unmarshal(reqRedis.Data, &req)
							if err != nil {
								continue
							}

							err = BranchUsecase.Create(req)
							if err != nil {
								continue
							}
						}
					}

					rdb.XAck(ctx, stream, groupName, message.ID)

				}
			}
		}

	}

}
