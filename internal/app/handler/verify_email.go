package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/labstack/echo/v4"
)

func VerifyEmail(
	verifyEmailLogic logic.VerifyEmailLogic,
) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.QueryParam("token")
		if token == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "token is required")
		}

		if err := verifyEmailLogic.VerifyEmail(context.Background(), token); err != nil {
			if errors.Is(err, logic.ErrVerificationTokenNotFound) {
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
		}

		return c.String(http.StatusOK, "OK")
	}
}
