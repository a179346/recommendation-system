package redishelper

import (
	"context"
	"time"

	"github.com/a179346/recommendation-system/internal/pkg/console"
	"github.com/redis/go-redis/v9"
)

func WaitForConnected(ctx context.Context, client *redis.Client) {
	for {
		select {
		case <-ctx.Done():
			return

		default:
			_, err := client.Ping(ctx).Result()
			if err == nil {
				return
			}
			console.Warnf("connecting to redis: %v", err)
			time.Sleep(2 * time.Second)
		}
	}
}
