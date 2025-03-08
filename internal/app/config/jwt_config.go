package config

import "github.com/a179346/recommendation-system/internal/pkg/envhelper"

type JwtConfig struct {
	Secret        []byte
	ExpireSeconds int
}

var jwtConfig JwtConfig

func init() {
	jwtConfig.Secret = []byte(envhelper.GetString("JWT_SECRET", "my1-jwt2-3secret"))
	jwtConfig.ExpireSeconds = envhelper.GetInt("JWT_EXPIRE_SECONDS", 3600)
}

func GetJwtConfig() JwtConfig {
	return jwtConfig
}
