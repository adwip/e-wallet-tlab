package utils

import (
	"errors"
	"fmt"

	"github.com/adwip/e-wallet-tlab/common-lib/stacktrace"
	"github.com/golang-jwt/jwt/v5"
)

type PayloadSchema struct {
	Name   string `json:"name"`
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(input PayloadSchema, jwtKey []byte) (string, error) {
	claims := PayloadSchema{
		UserID: input.UserID,
		Name:   input.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "e-wallet-tlab",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string, jwtKey []byte) (out *PayloadSchema, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &PayloadSchema{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
			return nil, stacktrace.CascadeWithClientMessage(err, stacktrace.FORBIDDEN, err.Error())
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, stacktrace.CascadeWithClientMessage(err, stacktrace.FORBIDDEN, err.Error())
	}
	if out, ok := token.Claims.(*PayloadSchema); ok && token.Valid {
		return out, nil
	}
	err = errors.New("invalid token")
	return nil, stacktrace.CascadeWithClientMessage(err, stacktrace.FORBIDDEN, err.Error())
}
