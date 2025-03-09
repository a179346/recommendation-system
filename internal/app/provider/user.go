package provider

import (
	"context"
	"database/sql"

	. "github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/table"
	"github.com/a179346/recommendation-system/internal/app/dto"
	. "github.com/go-jet/jet/v2/mysql"

	"github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/model"
)

type UserProvider struct {
	db *sql.DB
}

func NewUserProvider(db *sql.DB) UserProvider {
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

func (userProvider UserProvider) VerifyEmail(
	ctx context.Context,
	token string,
) (bool, error) {
	stmt := User.UPDATE().
		SET(
			User.Verified.SET(Bool(true)),
		).
		WHERE(
			User.Token.EQ(String(token)).
				AND(User.Verified.EQ(Bool(false))),
		)

	result, err := stmt.ExecContext(ctx, userProvider.db)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	return rowsAffected >= 1, nil
}

func (userProvider UserProvider) FindByEmail(ctx context.Context, email string) (dto.User, error) {
	stmt := SELECT(
		User.AllColumns,
	).FROM(
		User,
	).WHERE(
		User.Email.EQ(String(email)),
	).LIMIT(1)

	var dest model.User
	err := stmt.QueryContext(ctx, userProvider.db, &dest)
	return formatUser(dest), err
}

func formatUser(u model.User) dto.User {
	return dto.User{
		UserID:            u.UserID,
		Email:             u.Email,
		EncryptedPassword: u.EncryptedPassword,
		Token:             u.Token,
		Verified:          u.Verified,
		CreatedAt:         u.CreatedAt,
		UpdatedAt:         u.UpdatedAt,
	}
}
