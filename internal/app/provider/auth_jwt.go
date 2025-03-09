package provider

import (
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

type authJwtConfig struct {
	SigningMethod string
	NewClaimsFunc func() jwt.Claims
	SigningKey    []byte
}

func (authJwtProvider AuthJwtProvider) GetConfig() authJwtConfig {
	jwtConfig := config.GetJwtConfig()
	return authJwtConfig{
		SigningMethod: jwt.SigningMethodHS256.Name,
		NewClaimsFunc: func() jwt.Claims { return new(AuthJwtClaims) },
		SigningKey:    jwtConfig.Secret,
	}
}
