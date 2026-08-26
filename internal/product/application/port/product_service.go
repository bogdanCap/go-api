package port

import (
	"context"
	"github.com/bogdanCap/go-api/internal/product/domain"
)

type IProductService interface {
	List(ctx context.Context) (domain.Product, error)
	Create(
		ctx context.Context,
		product *domain.Product,
	) error
}
