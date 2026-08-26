package dto

import "time"

type ProductResponseDTO struct {
	/**
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`**/
	ArticleID int64   `json:"articleId"`
	Name      string  `json:"name"`
	Desc      string  `json:"desc"`
	//Price     float64 `json:"price"`
	//Status    Status
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
