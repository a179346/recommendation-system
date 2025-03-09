package logic

import (
	"context"
	"errors"

	"github.com/a179346/recommendation-system/internal/app/provider"
	"github.com/a179346/recommendation-system/internal/pkg/cryption"
	"github.com/go-jet/jet/v2/qrm"
)

type LoginLogic struct {
	userProvider    provider.UserProvider
	authJwtProvider provider.AuthJwtProvider
}

func NewLogin(
	userProvider provider.UserProvider,
	authJwtProvider provider.AuthJwtProvider,
) LoginLogic {
	return LoginLogic{
		userProvider:    userProvider,
		authJwtProvider: authJwtProvider,
	}
}

var ErrEmailNotFound = errors.New("Email not found")
var ErrEmailNotVerified = errors.New("Email not verified")
var ErrIncorrectPassword = errors.New("Incorrect password")

func (loginLogic LoginLogic) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	user, err := loginLogic.userProvider.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return "", ErrEmailNotFound
		}
		return "", err
	}

	if !user.Verified {
		return "", ErrEmailNotVerified
	}

	encryptedPassword := cryption.SHA256(password)
	if user.EncryptedPassword != encryptedPassword {
		return "", ErrIncorrectPassword
	}

	return loginLogic.authJwtProvider.Sign(user.UserID)
}
