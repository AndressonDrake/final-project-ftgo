package handler

import (
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
