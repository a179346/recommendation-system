package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/a179346/recommendation-system/internal/app/dto"
)

type ProductRedisProvider struct {
	client *redis.Client
}

func NewProductRedisProvider(client *redis.Client) ProductRedisProvider {
	return ProductRedisProvider{client: client}
}

func (productRedisProvider ProductRedisProvider) getRedisKey(cursor int, pageSize int) string {
	return fmt.Sprintf("product:%v:%v", cursor, pageSize)
}

func (productRedisProvider ProductRedisProvider) FindByCursorAndPageSize(
	ctx context.Context,
	cursor int,
	pageSize int,
) ([]dto.Product, error) {
	key := productRedisProvider.getRedisKey(cursor, pageSize)

	val, err := productRedisProvider.client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var products []dto.Product
	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (productRedisProvider ProductRedisProvider) SetByCursorAndPageSize(
	ctx context.Context,
	products []dto.Product,
	cursor int,
	pageSize int,
) error {
	key := productRedisProvider.getRedisKey(cursor, pageSize)

	val, err := json.Marshal(products)
	if err != nil {
		return err
	}

	return productRedisProvider.client.Set(ctx, key, string(val), 10*time.Minute).Err()
}
