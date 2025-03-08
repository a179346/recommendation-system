package config

import "github.com/a179346/recommendation-system/internal/pkg/envhelper"

type MigrationConfig struct {
	FolderPath string
	Verbose    bool
	Up         bool
}

var migrationConfig MigrationConfig

func init() {
	migrationConfig.FolderPath = envhelper.GetString("MIGRATION_FOLDER_PATH", "internal/app/migrations")
	migrationConfig.Verbose = envhelper.GetBool("MIGRATION_VERBOSE", true)
	migrationConfig.Up = envhelper.GetBool("MIGRATION_UP", true)
}

func GetMigrationConfig() MigrationConfig {
	return migrationConfig
}
