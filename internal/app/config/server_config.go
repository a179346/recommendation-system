package config

import "github.com/a179346/recommendation-system/internal/pkg/envhelper"

type ServerConfig struct {
	Port int
}

var serverConfig ServerConfig

func init() {
	serverConfig.Port = envhelper.GetInt("SERVER_PORT", 3000)
}

func GetServerConfig() ServerConfig {
	return serverConfig
}
