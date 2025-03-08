package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/labstack/echo/v4"
)

func RegisterUser(
	registerLogic logic.RegisterLogic,
) echo.HandlerFunc {
	type requestBody struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=16"`
	}

	specialCharacters := "()[]{}<>+-*/?,.:;\"'_\\|~`!@#$%^&="

	validatePassword := func(password string) bool {
		var hasUppercase bool
		var hasLowercase bool
		var hasSpecialCharacter bool

		for i := 0; i < len(password); i++ {
			b := password[i]
			if b >= 'a' && b <= 'z' {
				hasLowercase = true
			} else if b >= 'A' && b <= 'Z' {
				hasUppercase = true
			} else if strings.IndexByte(specialCharacters, b) != -1 {
				hasSpecialCharacter = true
			}
		}

		return hasUppercase && hasLowercase && hasSpecialCharacter
	}

	return func(c echo.Context) error {
		var body requestBody
		if err := c.Bind(&body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := c.Validate(&body); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if !validatePassword(body.Password) {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid password")
		}

		if err := registerLogic.RegisterUser(context.Background(), body.Email, body.Password); err != nil {
			if errors.Is(err, logic.ErrDuplicatedEmail) {
				return echo.NewHTTPError(http.StatusConflict, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
		}

		return c.String(http.StatusOK, "OK")
	}
}
