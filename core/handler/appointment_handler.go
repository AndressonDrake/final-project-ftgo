package handler

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type appointmentHandler struct {
	appointmentUsecase domain.AppointmentUsecase
}

func AppointmentHandler(appointmentUsecase domain.AppointmentUsecase) domain.AppointmentHandler {
	return &appointmentHandler{appointmentUsecase: appointmentUsecase}
}

func (h *appointmentHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetAppointment
	var responseErr model.ResponseError

	data, message, detail, err := h.appointmentUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
