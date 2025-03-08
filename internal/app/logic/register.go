package email_provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/a179346/recommendation-system/internal/app/config"
	"github.com/a179346/recommendation-system/internal/app/providers/email_provider"
	"github.com/a179346/recommendation-system/internal/app/providers/user_provider"
	"github.com/a179346/recommendation-system/internal/pkg/console"
	"github.com/a179346/recommendation-system/internal/pkg/cryption"
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type RegisterLogic struct {
	userProvider  user_provider.UserProvider
	emailProvider email_provider.EmailProvider
}

func NewRegister(
	userProvider user_provider.UserProvider,
	emailProvider email_provider.EmailProvider,
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
		if err, ok := err.(*mysql.MySQLError); ok {
			//TODO
			console.Infof("Number:%v SQLState:%v Message:%v", err.Number, err.SQLState, err.Message)
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
