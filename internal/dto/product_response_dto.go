package dto

import "time"

type ProductResponseDTO struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
