package migrations

import (
	"context"
	"database/sql"
	"strings"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up1741417443, Down1741417443)
}

func Up1741417443(ctx context.Context, tx *sql.Tx) error {
	query := []string{
		`CREATE TABLE user`,
		`(`,
		`		user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,`,
		`		email VARCHAR(255) NOT NULL UNIQUE,`,
		`		encrypted_password VARCHAR(255) NOT NULL,`,
		`		token VARCHAR(36) NOT NULL,`,
		`		verified TINYINT(1) NOT NULL DEFAULT 0,`,
		`		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,`,
		`		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,`,
		`		INDEX token (token) USING BTREE`,
		`);`,

		`CREATE TABLE product`,
		`(`,
		`		product_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,`,
		`		title VARCHAR(127) NOT NULL,`,
		`		price FLOAT(16,4) NOT NULL,`,
		`		description MEDIUMTEXT NOT NULL,`,
		`		category VARCHAR(63) NOT NULL`,
		`);`,
	}
	if _, err := tx.ExecContext(ctx, strings.Join(query, "\n")); err != nil {
		return err
	}

	return nil
}

func Down1741417443(ctx context.Context, tx *sql.Tx) error {
	query := []string{
		`DROP TABLE IF EXISTS product;`,
		`DROP TABLE IF EXISTS user;`,
	}
	if _, err := tx.ExecContext(ctx, strings.Join(query, "\n")); err != nil {
		return err
	}

	return nil
}
