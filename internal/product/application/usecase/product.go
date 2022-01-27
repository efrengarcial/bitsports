package usecase

import (
	"bitsports/internal/product/domain"
	"bitsports/internal/product/domain/repository"
	"bitsports/pkg/validate"
	"context"
)

type Product struct {
	repo repository.IProduct
}

// NewProduct will create new a Product object
func NewProduct(repo repository.IProduct) *Product {
	return &Product{
		repo,
	}
}

// Create insert information in the database
func (p Product) Create(ctx context.Context, prd domain.NewProduct) (*domain.Product, error) {
	if err := validate.Check(prd); err != nil {
		return nil, err
	}
	return p.repo.Create(ctx, prd)
}

// Update  modifies data about a Product.
func (p Product) Update(ctx context.Context, prd domain.UpdateProduct) (*domain.Product, error) {
	if err := validate.Check(prd); err != nil {
		return nil, err
	}
	return p.repo.Update(ctx, prd)
}

//Delete removes the product identified by a given ID.
func (p Product) Delete(ctx context.Context, productID int64) (err error) {
	return p.repo.Delete(ctx, productID)
}

// Query gets all Products from the database.
func (p Product) Query(ctx context.Context) ([]domain.Product, error) {
	return p.repo.Query(ctx)
}

// QueryByID finds the product identified by a given ID.
func (p Product) QueryByID(ctx context.Context, productID int64) (*domain.Product, error) {
	return p.repo.QueryByID(ctx, productID)
}
