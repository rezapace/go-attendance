package jwt_service

import (
	"presensee_project/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTService interface {
	GenerateToken(user *model.User) (string, error)
	GetClaims(c *echo.Context) jwt.MapClaims
}
