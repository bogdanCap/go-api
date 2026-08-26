package dto

type ProductDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

type ProductListFilterDTO struct {
	Name *string `validate:"omitempty,min=3"`
	Status *int `validate:"omitempty,gte=1,lte=4"`
	Limit int `validate:"omitempty,gte=0,lte=500"`
	Offset int `validate:"omitempty,gte=1"`
}
