package handler

import (
	"strconv"

	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type icd10Handler struct {
	icd10Usecase domain.ICD10Usecase
}

func ICD10Handler(icd10Usecase domain.ICD10Usecase) domain.ICD10Handler {
	return &icd10Handler{icd10Usecase: icd10Usecase}
}

func (h *icd10Handler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetICD10
	var responseErr model.ResponseError

	data, message, detail, err := h.icd10Usecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (h *icd10Handler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetICD10ByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := h.icd10Usecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
