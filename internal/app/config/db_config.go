package config

import "github.com/a179346/recommendation-system/internal/pkg/envhelper"

type DBConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
}

var dbConfig DBConfig

func init() {
	dbConfig.Host = envhelper.GetString("DB_HOST", "localhost")
	dbConfig.Port = envhelper.GetInt("DB_PORT", 3306)
	dbConfig.Database = envhelper.GetString("DB_DATABASE", "recommendation")
	dbConfig.User = envhelper.GetString("DB_USER", "recommendation-mysql-user")
	dbConfig.Password = envhelper.GetString("DB_PASSWORD", "recommendation-mysql-password")
}

func GetDBConfig() DBConfig {
	return dbConfig
}
