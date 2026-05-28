package handler

import (
	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type prescriptionHandler struct {
	prescriptionUsecase domain.PrescriptionUsecase
}

func PrescriptionHandler(prescriptionUsecase domain.PrescriptionUsecase) domain.PrescriptionHandler {
	return &prescriptionHandler{prescriptionUsecase: prescriptionUsecase}
}

func (h *prescriptionHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetPrescription
	var responseErr model.ResponseError

	data, message, detail, err := h.prescriptionUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
