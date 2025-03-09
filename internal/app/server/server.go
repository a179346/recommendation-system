package server

import (
	"database/sql"
	"net/http"

	"github.com/a179346/recommendation-system/internal/app/handler"
	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/a179346/recommendation-system/internal/app/provider"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	emailProvider := provider.NewEmailProvier()
	authJwtProvider := provider.NewAuthJwtProvider()

	registerLogic := logic.NewRegister(userProvider, emailProvider)
	VerifyEmailLogic := logic.NewVerifyEmail(userProvider)
	loginLogic := logic.NewLogin(userProvider, authJwtProvider)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))

	apiGroup := e.Group("/api")

	apiGroup.POST("/user/register", handler.RegisterUser(registerLogic))
	apiGroup.GET("/user/verify-email", handler.VerifyEmail(VerifyEmailLogic))
	apiGroup.POST("/user/login", handler.Login(loginLogic))

	return e
}
