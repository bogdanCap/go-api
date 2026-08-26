package port

import (
	"context"

	"github.com/bogdanCap/go-api/internal/product/domain"
	"github.com/bogdanCap/go-api/internal/product/dto"
)

type IProductRepository interface {
	//GetByID(ctx context.Context, id int64) (*product.Product, error)

	Create(
		ctx context.Context,
		product *domain.Product,
	) error
	List(
		ctx context.Context,
		filter dto.ProductListFilterDTO,
	) ([]domain.Product, error)
}
