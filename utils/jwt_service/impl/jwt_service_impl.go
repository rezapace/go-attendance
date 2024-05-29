package impl

import (
	"time"

	"presensee_project/utils/jwt_service"

	"presensee_project/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTServiceImpl struct {
	secretKey string
	exp       time.Duration
}

func NewJWTService(secretKey string, exp time.Duration) jwt_service.JWTService {
	return &JWTServiceImpl{
		secretKey: secretKey,
		exp:       exp,
	}
}

func (j *JWTServiceImpl) GenerateToken(user *model.User) (string, error) {
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(j.exp).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (*JWTServiceImpl) GetClaims(c *echo.Context) jwt.MapClaims {
	user := (*c).Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}
