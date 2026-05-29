package handler

import (
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gw-heatlcare.com/helper"
	"gw-heatlcare.com/model"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var response model.ResponseErrorPost
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			response.Message = "nil token"
			return c.JSON(403, response)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := helper.VerifyToken(tokenString)

		if err != nil || !token.Valid {
			response.Message = err.Error()
			return c.JSON(403, response)
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := claims["user_id"].(float64)

		c.Set("user_id", userID)

		return next(c)

	}
}
