package handler

import (
	"strconv"

	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type healthNewsHandler struct {
	healthNewsUsecase domain.HealthNewsUsecase
}

func HealthNewsHandler(healthNewsUsecase domain.HealthNewsUsecase) domain.HealthNewsHandler {
	return &healthNewsHandler{healthNewsUsecase: healthNewsUsecase}
}

func (h *healthNewsHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetHealthNews
	var responseErr model.ResponseError

	data, message, detail, err := h.healthNewsUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (h *healthNewsHandler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetHealthNewsByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := h.healthNewsUsecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
