package handler

import (
	"net/http"

	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/labstack/echo/v4"
)

func RegisterUser(
	registerLogic logic.RegisterLogic,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO
		return c.String(http.StatusOK, "Hello, World!")
	}
}
