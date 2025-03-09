package logic

import (
	"context"

	"github.com/a179346/recommendation-system/internal/app/dto"
	"github.com/a179346/recommendation-system/internal/app/provider"
)

type GetRecommendationLogic struct {
	productDbProvider    provider.ProductDbProvider
	productRedisProvider provider.ProductRedisProvider
}

func NewGetRecommendation(
	productDbProvider provider.ProductDbProvider,
	productRedisProvider provider.ProductRedisProvider,
) GetRecommendationLogic {
	return GetRecommendationLogic{
		productDbProvider:    productDbProvider,
		productRedisProvider: productRedisProvider,
	}
}

func (getRecommendationLogic GetRecommendationLogic) GetRecommendation(
	ctx context.Context,
	cursor int,
	pageSize int,
) ([]dto.Product, error) {
	if products, err := getRecommendationLogic.productRedisProvider.FindByCursorAndPageSize(ctx, cursor, pageSize); err == nil {
		return products, nil
	}

	products, err := getRecommendationLogic.productDbProvider.FindByCursorAndPageSize(ctx, cursor, pageSize)
	if err != nil {
		return nil, err
	}

	if err := getRecommendationLogic.productRedisProvider.SetByCursorAndPageSize(ctx, products, cursor, pageSize); err != nil {
		return nil, err
	}

	return products, nil
}
