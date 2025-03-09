package provider

import (
	"errors"
	"fmt"
	"time"

	"github.com/a179346/recommendation-system/internal/app/config"
	"github.com/golang-jwt/jwt/v5"
)

type AuthJwtClaims struct {
	ID int32 `json:"id"`
	jwt.RegisteredClaims
}

type AuthJwtProvider struct{}

func NewAuthJwtProvider() AuthJwtProvider {
	return AuthJwtProvider{}
}

func (authJwtProvider AuthJwtProvider) Sign(id int32) (string, error) {
	jwtConfig := config.GetJwtConfig()

	claims := AuthJwtClaims{
		id,
		jwt.RegisteredClaims{
			Issuer:    "recommendation",
			Subject:   "auth",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now().Add(-10 * time.Minute)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.ExpireSeconds) * time.Second)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtConfig.Secret)
}

// TODO remove ?
func (authJwtProvider AuthJwtProvider) Parse(tokenString string) (*AuthJwtClaims, error) {
	jwtConfig := config.GetJwtConfig()

	token, err := jwt.ParseWithClaims(tokenString, &AuthJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtConfig.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AuthJwtClaims)
	if !ok {
		return nil, errors.New("unknown jwt payload")
	}
	return claims, nil
}
