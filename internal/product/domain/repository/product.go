package repository

import (
	"bitsports/internal/product/domain"
	"context"
)

// IProduct  is interface of Product repository
type IProduct interface {
	// Create insert information in the database
	Create(ctx context.Context, prd domain.NewProduct) (*domain.Product, error)
	// Update  modifies data about a Product.
	Update(ctx context.Context, prd domain.UpdateProduct) (*domain.Product, error)
	//Delete removes the product identified by a given ID.
	Delete(ctx context.Context, productID int64) (err error)
	// Query gets all Products from the database.
	Query(ctx context.Context) ([]domain.Product, error)
	// QueryByID finds the product identified by a given ID.
	QueryByID(ctx context.Context, productID int64) (*domain.Product, error)
}
