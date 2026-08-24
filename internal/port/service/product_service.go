package service

import (
	"context"
	"github.com/bogdanCap/go-api/internal/domain/product"
)

type IProductService interface {
	List(ctx context.Context) (product.Product, error)
	Create(
		ctx context.Context,
		product *product.Product,
	) error
}
