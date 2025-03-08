package user_provider

import (
	"context"
	"database/sql"

	. "github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/table"
	. "github.com/go-jet/jet/v2/mysql"

	"github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/model"
)

type UserProvider struct {
	db *sql.DB
}

func New(db *sql.DB) UserProvider {
	return UserProvider{db: db}
}

func (userProvider UserProvider) CreateUser(
	ctx context.Context,
	email string,
	encryptedPassword string,
	token string,
) error {
	newUser := model.User{
		Email:             email,
		EncryptedPassword: encryptedPassword,
		Token:             token,
	}

	columnList := ColumnList{User.Email, User.EncryptedPassword, User.Token}
	stmt := User.INSERT(columnList).MODEL(newUser)

	_, err := stmt.ExecContext(ctx, userProvider.db)
	return err
}
