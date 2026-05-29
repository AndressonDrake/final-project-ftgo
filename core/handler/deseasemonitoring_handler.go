package handler

import (
	"strconv"

	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type diseaseMonitoringHandler struct {
	diseaseMonitoringUsecase domain.DiseaseMonitoringUsecase
}

func DiseaseMonitoringHandler(diseaseMonitoringUsecase domain.DiseaseMonitoringUsecase) domain.DiseaseMonitoringHandler {
	return &diseaseMonitoringHandler{diseaseMonitoringUsecase: diseaseMonitoringUsecase}
}

func (h *diseaseMonitoringHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetDiseaseMonitoring
	var responseErr model.ResponseError

	data, message, detail, err := h.diseaseMonitoringUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (h *diseaseMonitoringHandler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetDiseaseMonitoringByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := h.diseaseMonitoringUsecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
