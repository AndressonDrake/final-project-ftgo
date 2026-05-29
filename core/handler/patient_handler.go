package handler

import (
	"strconv"

	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type patientHandler struct {
	patientUsecase domain.PatientUsecase
}

func PatientHandler(patientUsecase domain.PatientUsecase) domain.PatientHandler {
	return &patientHandler{patientUsecase: patientUsecase}
}

func (h *patientHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetPatient
	var responseErr model.ResponseError

	data, message, detail, err := h.patientUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (m *patientHandler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetPatientByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := m.patientUsecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
