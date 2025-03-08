package dbhelper

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/a179346/recommendation-system/internal/app/config"
	"github.com/a179346/recommendation-system/internal/pkg/console"
	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	dbConfig := config.GetDBConfig()

	databaseURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=true&multiStatements=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)

	return sql.Open("mysql", databaseURL)
}

func WaitFor(ctx context.Context, db *sql.DB) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			_, err := db.Query("SELECT 1")
			if err == nil {
				return
			}
			console.Warnf("connecting to database: %v", err)
			time.Sleep(2 * time.Second)
		}
	}
}
