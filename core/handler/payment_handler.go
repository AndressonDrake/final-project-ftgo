package handler

import (
	"strconv"

	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	paymentUsecase domain.PaymentUsecase
}

func PaymentHandler(paymentUsecase domain.PaymentUsecase) domain.PaymentHandler {
	return &paymentHandler{paymentUsecase: paymentUsecase}
}

func (h *paymentHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetPayment
	var responseErr model.ResponseError

	data, message, detail, err := h.paymentUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (m *paymentHandler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetPaymentByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := m.paymentUsecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
