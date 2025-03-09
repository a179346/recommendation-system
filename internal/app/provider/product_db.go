package provider

import (
	"context"
	"database/sql"
	"time"

	. "github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/table"
	"github.com/a179346/recommendation-system/internal/app/dto"
	"github.com/a179346/recommendation-system/internal/pkg/slicehelper"
	. "github.com/go-jet/jet/v2/mysql"

	"github.com/a179346/recommendation-system/internal/app/database/.jet_gen/recommendation/model"
)

type ProductDbProvider struct {
	db *sql.DB
}

func NewProductDbProvider(db *sql.DB) ProductDbProvider {
	return ProductDbProvider{db: db}
}

func (productDbProvider ProductDbProvider) FindByCursorAndPageSize(ctx context.Context, cursor int, pageSize int) ([]dto.Product, error) {
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

	// Simulate a 2-second delay to mimic the time taken for a database query.
	time.Sleep(2 * time.Second)

	return slicehelper.Map(dest, formatProduct), err
}

func formatProduct(p model.Product) dto.Product {
	return dto.Product{
		ProductID:   p.ProductID,
		Title:       p.Title,
		Price:       p.Price,
		Description: p.Description,
		Category:    p.Category,
	}
}
