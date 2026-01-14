package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService struct {
	secretKey []byte
}

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: []byte(os.Getenv("JWT_SECRET")),
	}
}

func (j *JWTService) GenerateToken(userID uint, email string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secretKey)
}
