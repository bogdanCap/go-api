package port

import (
	"context"
	"github.com/bogdanCap/go-api/internal/product/domain"
	"github.com/bogdanCap/go-api/internal/product/dto"
)

type IProductService interface {
	List(ctx context.Context, filter dto.ProductListFilterDTO) ([]domain.Product, error)
	Create(
		ctx context.Context,
		product *domain.Product,
	) error
}
