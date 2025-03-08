package server

import (
	"database/sql"
	"net/http"

	"github.com/a179346/recommendation-system/internal/app/handler"
	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/a179346/recommendation-system/internal/app/providers/email_provider"
	"github.com/a179346/recommendation-system/internal/app/providers/user_provider"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetServer(db *sql.DB) *echo.Echo {
	userProvider := user_provider.New(db)
	emailProvider := email_provider.New()

	registerLogic := logic.NewRegister(userProvider, emailProvider)

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))

	apiGroup := e.Group("/api")

	apiGroup.POST("/user/register", handler.RegisterUser(registerLogic))

	return e
}
