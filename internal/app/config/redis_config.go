package config

import "github.com/a179346/recommendation-system/internal/pkg/envhelper"

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

var redisConfig RedisConfig

func init() {
	redisConfig.Host = envhelper.GetString("REDIS_HOST", "localhost")
	redisConfig.Port = envhelper.GetInt("REDIS_PORT", 6379)
	redisConfig.Password = envhelper.GetString("REDIS_PASSWORD", "recommendation-redis-pass")
	redisConfig.DB = envhelper.GetInt("REDIS_DB", 0)
	redisConfig.PoolSize = envhelper.GetInt("REDIS_POOL_SIZE", 10)
}

func GetRedisConfig() RedisConfig {
	return redisConfig
}
