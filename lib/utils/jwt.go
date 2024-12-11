package utils

import (
	"errors"

	"ardeolib.sapions.com/models"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	models.User
	jwt.RegisteredClaims
}

type JWTUtil interface {
	GetSecretKey() string
	GenerateToken(claims *UserClaims) (string, error)
	ValidateToken(tokenString string) (userclaims *UserClaims, err error)
}

type jwtUtil struct {
	secretKey string
}

func NewJWTUtil(secretKey string) JWTUtil {
	return &jwtUtil{secretKey: secretKey}
}

func (j *jwtUtil) GetSecretKey() string {
	return j.secretKey
}

func (j *jwtUtil) GenerateToken(claims *UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtUtil) ValidateToken(tokenString string) (*UserClaims, error) {
	var userclaims UserClaims
	var err error
	_, err = jwt.ParseWithClaims(tokenString, &userclaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return &userclaims, nil
}
