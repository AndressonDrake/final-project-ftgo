package handler

import (
	"github.com/labstack/echo/v4"
	"gw-heatlcare.com/domain"
	"gw-heatlcare.com/model"
)

type postHandler struct {
	postUsecase domain.PostUsecase
}

func PostHandler(postUsecae domain.PostUsecase) domain.PostHandler {
	return &postHandler{postUsecase: postUsecae}
}

func (p *postHandler) Post(c echo.Context) (err error) {

	var request model.RequestGeneral
	var response model.ResponseSuccessPost
	var responseErr model.ResponseErrorPost

	if err := c.Bind(&request); err != nil {
		responseErr.Message = "bad request"
		responseErr.Detail = err.Error()
		return c.JSON(400, responseErr)
	}

	message, detail, err := p.postUsecase.Post(request)

	responseErr.Detail = detail
	responseErr.Message = message

	if err != nil {
		return c.JSON(500, responseErr)
	}

	response.Message = message

	return c.JSON(200, response)
}
