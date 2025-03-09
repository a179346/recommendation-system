package server

import (
	"database/sql"
	"net/http"

	"github.com/a179346/recommendation-system/internal/app/handler"
	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/a179346/recommendation-system/internal/app/middleware"
	"github.com/a179346/recommendation-system/internal/app/provider"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	echomiddlewawre "github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetServer(db *sql.DB) *echo.Echo {
	userProvider := provider.NewUserProvider(db)
	productDbProvider := provider.NewProductDbProvider(db)
	emailProvider := provider.NewEmailProvier()
	authJwtProvider := provider.NewAuthJwtProvider()

	registerLogic := logic.NewRegister(userProvider, emailProvider)
	VerifyEmailLogic := logic.NewVerifyEmail(userProvider)
	loginLogic := logic.NewLogin(userProvider, authJwtProvider)
	getRecommendationLogic := logic.NewGetRecommendation(productDbProvider)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(echomiddlewawre.Recover())
	e.Use(echomiddlewawre.BodyLimit("2M"))

	apiGroup := e.Group("/api")

	apiGroup.POST("/user/register", handler.RegisterUser(registerLogic))
	apiGroup.GET("/user/verify-email", handler.VerifyEmail(VerifyEmailLogic))
	apiGroup.POST("/user/login", handler.Login(loginLogic))

	authedGroup := apiGroup.Group("/authed")
	authedGroup.Use(middleware.JwtAuth(authJwtProvider))

	authedGroup.GET("/recommendation", handler.GetRecommendation(getRecommendationLogic))

	return e
}
