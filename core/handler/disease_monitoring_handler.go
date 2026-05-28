package handler

import (
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
