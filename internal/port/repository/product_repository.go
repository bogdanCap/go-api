package repository

import (
	"context"
	"test/internal/domain/product"
)

type IProductRepository interface {
	GetByID(ctx context.Context, id int64) (*product.Product, error)
}
