package handler

import (
	"github.com/labstack/echo/v4"
	"gw-heatlcare.com/domain"
	"gw-heatlcare.com/model"
)

type coreHandler struct {
	coreUsecase domain.CoreUsecase
}

func CoreHandler(coreUsecase domain.CoreUsecase) domain.CoreHandler {
	return &coreHandler{coreUsecase: coreUsecase}
}

func (c *coreHandler) GetCore(e echo.Context) (err error) {
	var responseSuccess model.ResponseGetCore
	var responseErr model.ResponseErrorPost

	pathUrl := e.QueryParam("url")

	data, message, detail, err := c.coreUsecase.GetCore(pathUrl)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return e.JSON(500, responseErr)
	}

	responseSuccess.Message = message
	responseSuccess.Data = data

	return e.JSON(200, responseSuccess)
}
