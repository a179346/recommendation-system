package email_provider

import (
	"context"

	"github.com/a179346/recommendation-system/internal/pkg/console"
)

type EmailProvider struct{}

func New() EmailProvider {
	return EmailProvider{}
}

func (emailProvider EmailProvider) SendEmailVerification(
	ctx context.Context,
	email string,
	verificationLink string,
) error {
	console.Infof("Sended email to %s with link: %s", email, verificationLink)
	return nil
}
