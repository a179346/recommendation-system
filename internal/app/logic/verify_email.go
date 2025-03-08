package logic

import (
	"context"
	"errors"

	"github.com/a179346/recommendation-system/internal/app/provider"
)

type VerifyEmailLogic struct {
	userProvider provider.UserProvider
}

func NewVerifyEmail(
	userProvider provider.UserProvider,
) VerifyEmailLogic {
	return VerifyEmailLogic{
		userProvider: userProvider,
	}
}

var ErrVerificationTokenNotFound = errors.New("Token isn't found or has been used")

func (verifyEmailLogic VerifyEmailLogic) VerifyEmail(
	ctx context.Context,
	token string,
) error {
	updated, err := verifyEmailLogic.userProvider.VerifyEmail(ctx, token)
	if err != nil {
		return err
	}
	if !updated {
		return ErrVerificationTokenNotFound
	}

	return nil
}
