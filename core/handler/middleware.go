package handler

import (

	"core-healtcare.com/model"
	"github.com/labstack/echo/v4"
)

var (
	API_KEY string
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var response model.ResponseError
		authHeader := c.Request().Header.Get("api-key")

		if authHeader == "" {
			response.Message = "nil token"
			return c.JSON(403, response)
		}

		if authHeader != API_KEY {
			response.Message = "invalid api key"
			return c.JSON(403, response)
		}

		return next(c)

	}
}
