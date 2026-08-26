package mapper

import (
	"github.com/bogdanCap/go-api/internal/product/domain"
	"github.com/bogdanCap/go-api/internal/product/dto"
)

func ProductToDomain(
	dto dto.ProductDTO,
) *domain.Product {

	return domain.NewProduct(
		dto.ID,
		dto.Name,
		dto.Price,
	)
}

func ProductToResponse(
	p domain.Product,
) dto.ProductResponseDTO {

	return dto.ProductResponseDTO{
		ID:        p.ID,
		Price:     p.Price,
		Name:      p.Name,
		Status:    p.Status.String(),
		CreatedAt: p.CreatedAt,
	}
}

func ProductsToResponse(
	products []domain.Product,
) []dto.ProductResponseDTO {

	result := make(
		[]dto.ProductResponseDTO,
		0,
		len(products),
	)

	for _, item := range products {

		result = append(
			result,
			ProductToResponse(item),
		)
	}

	return result
}
