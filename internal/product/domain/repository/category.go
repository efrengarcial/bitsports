package repository

import (
	"bitsports/internal/product/domain"
	"context"
)

// ICategory  is interface of Category repository
type ICategory interface {
	// Create insert information in the database
	Create(ctx context.Context, cat domain.NewCategory) (*domain.Category, error)
	// Update  modifies data about a Category.
	Update(ctx context.Context, cat domain.UpdateCategory) (*domain.Category, error)
	//Delete removes the category identified by a given ID.
	Delete(ctx context.Context, categoryID int64) (err error)
	// Query gets all Categories from the database.
	Query(ctx context.Context) ([]domain.Category, error)
	// QueryByID finds the category identified by a given ID.
	QueryByID(ctx context.Context, categoryID int64) (*domain.Category, error)
}

