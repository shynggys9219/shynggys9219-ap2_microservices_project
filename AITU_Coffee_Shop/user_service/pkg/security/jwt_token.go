package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{secretKey: secretKey}
}

func (j *JWTManager) GenerateAccessToken(userID uint64, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) GenerateRefreshToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTManager) Verify(tokenStr string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	return claims, err
}
