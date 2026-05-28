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

func (h *medicalRecordHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetMedicalRecord
	var responseErr model.ResponseError

	data, message, detail, err := h.medicalRecordUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
