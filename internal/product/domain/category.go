package domain

import "time"

// Category represents a product category.
type Category struct {
	ID        int64     `json:"id"`        // Unique identifier.
	Code      string    `json:"code"`      // Display code of the category.
	Name      string    `json:"name"`      // Display name of the category.
	CreatedAt time.Time `json:"createdAt"` //// CreatedAt holds the value of the "created_at" field.
	UpdatedAt time.Time `json:"updatedAt"` // UpdatedAt holds the value of the "updated_at" field.
}

// NewCategory is what we require from clients when adding a Category.
type NewCategory struct {
	Name string `json:"name" validate:"required,alpha,lte=100"`
	Code string `json:"code" validate:"required,alpha,lte=20"`
}

type UpdateCategory struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"  validate:"omitempty,alpha,lte=100"`
	Code *string `json:"code" validate:"omitempty,alpha,lte=20"`
}
