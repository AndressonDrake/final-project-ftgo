package handler

import (
	"strconv"

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

func (m *medicalRecordHandler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetMedicalRecordByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := m.medicalRecordUsecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
