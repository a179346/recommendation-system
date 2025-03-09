package middleware

import (
	"github.com/a179346/recommendation-system/internal/app/provider"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JwtAuth(authJwtProvider provider.AuthJwtProvider) echo.MiddlewareFunc {
	authJwtConfig := authJwtProvider.GetConfig()
	config := echojwt.Config{
		SigningMethod: authJwtConfig.SigningMethod,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return authJwtConfig.NewClaimsFunc()
		},
		SigningKey: authJwtConfig.SigningKey,
	}
	return echojwt.WithConfig(config)
}
