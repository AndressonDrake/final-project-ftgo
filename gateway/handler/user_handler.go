package handler

import (
	"github.com/labstack/echo/v4"
	"gw-heatlcare.com/domain"
	"gw-heatlcare.com/model"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func UserHandler(userUsecase domain.UserUsecase) domain.UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (u *userHandler) Login(c echo.Context) (err error) {
	var request model.ReqeustLoginUser
	var responseOK model.ResponseSuccessLoginUser
	var responseErr model.ResponseErrorPost

	if err := c.Bind(&request); err != nil {
		responseErr.Message = "bad request"
		responseErr.Detail = err.Error()
		return c.JSON(400, responseErr)
	}

	token, message, detail, err := u.userUsecase.Login(request)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	responseOK.Message = message
	responseOK.Token = token

	return c.JSON(200, responseOK)

}
