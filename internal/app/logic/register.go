package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/a179346/recommendation-system/internal/app/config"
	"github.com/a179346/recommendation-system/internal/app/provider"
	"github.com/a179346/recommendation-system/internal/pkg/cryption"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type RegisterLogic struct {
	userProvider  provider.UserProvider
	emailProvider provider.EmailProvider
}

func NewRegister(
	userProvider provider.UserProvider,
	emailProvider provider.EmailProvider,
) RegisterLogic {
	return RegisterLogic{
		userProvider:  userProvider,
		emailProvider: emailProvider,
	}
}

var ErrDuplicatedEmail = errors.New("email has been taken")

func (registerLogic RegisterLogic) RegisterUser(
	ctx context.Context,
	email string,
	password string,
) error {
	encryptedPassword := cryption.SHA256(password)
	token := uuid.New().String()

	if err := registerLogic.userProvider.CreateUser(ctx, email, encryptedPassword, token); err != nil {
		// 1062: Duplicate entry
		if err, ok := err.(*mysql.MySQLError); ok && err.Number == 1062 {
			return ErrDuplicatedEmail
		}
		return err
	}

	verificationLink := fmt.Sprintf(
		"http://localhost:%v/api/user/verify-email?token=%s",
		config.GetServerConfig().Port,
		token,
	)
	if err := registerLogic.emailProvider.SendEmailVerification(ctx, email, verificationLink); err != nil {
		return err
	}

	return nil
}
