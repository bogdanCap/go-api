package domain

import "time"

type Product struct {
	ArticleID int64 
	Name      string  
	Desc      string 
	//Price     float64 `json:"price"`
	//Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

/*
func NewProduct(id string, name string, price float64) *Product {
	return &Product{
		ID:        id,
		Name:      name,
		Price:     price,
		Status:    StatusNew,
		CreatedAt: time.Now(),
	}
}*/

/*
func (p *Product) CanCancel() bool {
	return p.Status == StatusNew ||
		p.Status == StatusProcessing
}

func (p *Product) Cancel() error {
	if !p.CanCancel() {
		return ErrProductCannotBeCancelled
	}

	p.Status = StatusCancelled

	return nil
}
*/