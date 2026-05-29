package handler

import (
	"strconv"

	"core-healtcare.com/domain"
	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

type medicineHandler struct {
	medicineUsecase domain.MedicineUsecase
}

func MedicineHandler(medicineUsecase domain.MedicineUsecase) domain.MedicineHandler {
	return &medicineHandler{medicineUsecase: medicineUsecase}
}

func (m *medicineHandler) Get(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetMedicine
	var responseErr model.ResponseError

	data, message, detail, err := m.medicineUsecase.Get()

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}

func (m *medicineHandler) GetByID(c echo.Context) (err error) {
	var responseOK model.ResponseSuccessGetMedicineByID
	var responseErr model.ResponseError

	idStr := c.Param("id")
	id, errConv := strconv.Atoi(idStr)
	if errConv != nil {
		responseErr.Message = "invalid id"
		responseErr.Detail = errConv.Error()
		return c.JSON(400, responseErr)
	}

	data, message, detail, err := m.medicineUsecase.GetByID(id)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(404, responseErr)
	}

	responseOK.Message = message
	responseOK.Data = data

	return c.JSON(200, responseOK)
}
