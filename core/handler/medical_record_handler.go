package handler

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type medicalRecordHandler struct {
	medicalRecordUsecase domain.MedicalRecordUsecase
}

func MedicalRecordHandler(medicalRecordUsecase domain.MedicalRecordUsecase) domain.MedicalRecordHandler {
	return &medicalRecordHandler{medicalRecordUsecase: medicalRecordUsecase}
}

func (mrh *medicalRecordHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetMedicalRecord
	var responseErr model.ResponseError

	data, message, detail, err := mrh.medicalRecordUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (mrh *medicalRecordHandler) Create(c echo.Context) (err error) {
	var request model.CreateMedicalRecord
	var responseErr model.ResponseError

	if err != nil {
		return c.JSON(500, responseErr)
	}

	err = mrh.medicalRecordUsecase.Create(request)
	if err != nil {
		return c.JSON(500, responseErr)
	}

	return c.JSON(201, map[string]string{"message": "success create medical record"})
}

func (mrh *medicalRecordHandler) Update(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{"message": "success update medical record"})
}

func (mrh *medicalRecordHandler) Delete(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{"message": "success delete medical record"})
}
