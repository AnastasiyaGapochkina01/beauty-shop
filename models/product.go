package models

import "time"

type Product struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Price        float64   `json:"price"`
	Brand        string    `json:"brand"`
	Category     string    `json:"category"`
	StockQuantity int      `json:"stock_quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Price        float64 `json:"price" validate:"required,gt=0"`
	Brand        string  `json:"brand" validate:"required"`
	Category     string  `json:"category" validate:"required"`
	StockQuantity int    `json:"stock_quantity" validate:"gte=0"`
}

type UpdateProductRequest struct {
	Name         *string  `json:"name"`
	Description  *string  `json:"description"`
	Price        *float64 `json:"price"`
	Brand        *string  `json:"brand"`
	Category     *string  `json:"category"`
	StockQuantity *int    `json:"stock_quantity"`
}
