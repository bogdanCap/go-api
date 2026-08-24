package repository

import (
	"context"
	"github.com/bogdanCap/go-api/internal/domain/product"
)

type IProductRepository interface {
	//GetByID(ctx context.Context, id int64) (*product.Product, error)

	Create(
		ctx context.Context,
		product *product.Product,
	) error
	List(
		ctx context.Context,
		limit int,
		offset int,
	) ([]*product.Product, error)
}
