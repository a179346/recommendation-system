package provider

import (
	"context"
	"database/sql"

	. "github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/table"
	. "github.com/go-jet/jet/v2/mysql"

	"github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/model"
)

type ProductDbProvider struct {
	db *sql.DB
}

func NewProductDbProvider(db *sql.DB) ProductDbProvider {
	return ProductDbProvider{db: db}
}

func (productDbProvider ProductDbProvider) FindByCursorAndPageSize(ctx context.Context, cursor int, pageSize int) ([]model.Product, error) {
	stmt := SELECT(
		Product.AllColumns,
	).FROM(
		Product,
	).WHERE(
		Product.ProductID.GT(Int(int64(cursor))),
	).ORDER_BY(
		Product.ProductID.ASC(),
	).LIMIT(
		int64(pageSize),
	)

	var dest []model.Product
	err := stmt.QueryContext(ctx, productDbProvider.db, &dest)
	return dest, err
}
