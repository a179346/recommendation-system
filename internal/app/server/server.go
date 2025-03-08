package server

import (
	"database/sql"

	"github.com/a179346/recommendation-system/internal/app/handler"
	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/a179346/recommendation-system/internal/app/providers/email_provider"
	"github.com/a179346/recommendation-system/internal/app/providers/user_provider"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetServer(db *sql.DB) *echo.Echo {
	userProvider := user_provider.New(db)
	emailProvider := email_provider.New()

	registerLogic := logic.NewRegister(userProvider, emailProvider)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.BodyLimit("2M"))

	apiGroup := e.Group("/api")

	apiGroup.GET("/user/register", handler.RegisterUser(registerLogic))

	return e
}
