package dto

type ProductDTO struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ProductListFilterDTO struct {
	Name *string `validate:"omitempty,min=3"`
	Age  *int    `validate:"omitempty,gte=18,lte=100"`
}
