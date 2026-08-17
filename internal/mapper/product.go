package mapper

import (
	"test/internal/domain/product"
	"test/internal/dto"
)

func ProductToDomain(
	dto dto.ProductDTO,
) *product.Product {

	return product.NewProduct(
		dto.Id,
		dto.Name,
		dto.Price,
	)
}

func ProductToResponse(
	p product.Product,
) dto.ProductResponseDTO {

	return dto.ProductResponseDTO{
		Id:        p.Id,
		Price:     p.Price,
		Name:      p.Name,
		Status:    p.Status.String(),
		CreatedAt: p.CreatedAt,
	}
}

func ProductsToResponse(
	products []product.Product,
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
