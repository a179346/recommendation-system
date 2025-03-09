package handler

import (
	"errors"
	"net/http"

	"github.com/a179346/recommendation-system/internal/app/logic"
	"github.com/a179346/recommendation-system/internal/pkg/set"
	"github.com/labstack/echo/v4"
)

func Login(
	loginLogic logic.LoginLogic,
) echo.HandlerFunc {
	type requestBody struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6,max=16"`
	}

	specialCharacters := "()[]{}<>+-*/?,.:;\"'_\\|~`!@#$%^&="
	specialCharacterSet := set.New[byte]()
	for i := 0; i < len(specialCharacters); i++ {
		specialCharacterSet.Add(specialCharacters[i])
	}

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
			} else if specialCharacterSet.Has(b) {
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

		token, err := loginLogic.Login(c.Request().Context(), body.Email, body.Password)
		if err != nil {
			if errors.Is(err, logic.ErrEmailNotFound) || errors.Is(err, logic.ErrIncorrectPassword) {
				return echo.NewHTTPError(http.StatusNotFound, "Email not found or incorrect password")
			}
			if errors.Is(err, logic.ErrEmailNotVerified) {
				return echo.NewHTTPError(http.StatusForbidden, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": token,
		})
	}
}
