package logic

import (
	"context"

	"github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/model"
	"github.com/a179346/recommendation-system/internal/app/provider"
)

type GetRecommendationLogic struct {
	productDbProvider provider.ProductDbProvider
}

func NewGetRecommendation(
	productDbProvider provider.ProductDbProvider,
) GetRecommendationLogic {
	return GetRecommendationLogic{
		productDbProvider: productDbProvider,
	}
}

// TODO redis
func (getRecommendationLogic GetRecommendationLogic) GetRecommendation(
	ctx context.Context,
	cursor int,
	pageSize int,
) ([]model.Product, error) {
	return getRecommendationLogic.productDbProvider.FindByCursorAndPageSize(ctx, cursor, pageSize)
}
