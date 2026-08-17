package service

import (
	"context"
	"test/internal/domain/product"
)

type IProductService interface {
	List(ctx context.Context) (product.Product, error)
	Create(
		ctx context.Context,
		product *product.Product,
	) error
}
