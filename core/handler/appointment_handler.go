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

func (ah *appointmentHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetAppointment
	var responseErr model.ResponseError

	data, message, detail, err := ah.appointmentUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (ah *appointmentHandler) Create(c echo.Context) (err error) {
	var request model.CreateAppointment
	var responseErr model.ResponseError

	if err != nil {
		return c.JSON(400, responseErr)
	}

	err = ah.appointmentUsecase.Create(request)
	if err != nil {
		return c.JSON(500, responseErr)
	}

	return c.JSON(201, map[string]string{"message": "success create appointment"})
}

func (ah *appointmentHandler) Update(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{"message": "success update appointment"})
}

func (ah *appointmentHandler) Delete(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{"message": "success delete appointment"})
}
