package domain

import (
	"time"
)

// Product represents an individual product.
type Product struct {
	ID        int64     `json:"id"`        // Unique identifier.
	Name      string    `json:"name"`      // Display name of the product.
	Price     float64   `json:"price"`     // Price for one item in cents.
	Quantity  int64     `json:"quantity"`  // Original number of items available.
	Category  Category  `json:"category"`  // Product Category.
	CreatedAt time.Time `json:"createdAt"` //// CreatedAt holds the value of the "created_at" field.
	UpdatedAt time.Time `json:"updatedAt"` // UpdatedAt holds the value of the "updated_at" field.
}

// NewProduct is what we require from clients when adding a Product.
type NewProduct struct {
	Name       string  `json:"name" validate:"required,alpha,lte=100"`
	Price      float64 `json:"price" validate:"required,lte=9999999"`
	Quantity   int64   `json:"quantity" validate:"required,gte=1"`
	CategoryID int64   `json:"categoryId" validate:"required"`
}

type UpdateProduct struct {
	ID         int64    `json:"id"`
	Name       *string  `json:"name"  validate:"omitempty,alpha,lte=100"`
	Price      *float64 `json:"price" validate:"omitempty,lte=9999999"`
	Quantity   *int64   `json:"quantity" validate:"omitempty,gte=1"`
	CategoryID *int64   `json:"categoryId"`
}
