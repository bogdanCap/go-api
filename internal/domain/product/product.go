package product

import "time"

type Product struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Status    Status
	CreatedAt time.Time
}

func NewProduct(id string, name string, price float64) *Product {
	return &Product{
		Id:        id,
		Name:      name,
		Price:     price,
		Status:    StatusNew,
		CreatedAt: time.Now(),
	}
}

func (o *Product) CanCancel() bool {
	return o.Status == StatusNew ||
		o.Status == StatusProcessing
}

func (o *Product) Cancel() error {
	if !o.CanCancel() {
		return ErrProductCannotBeCancelled
	}

	o.Status = StatusCancelled

	return nil
}
