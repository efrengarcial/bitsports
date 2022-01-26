package usecase

import (
	"bitsports/internal/product/domain"
	"bitsports/internal/product/domain/repository"
	"context"
)

type Category struct {
	repo repository.ICategory
}

// NewCategory will create new a Category object
func NewCategory(repo repository.ICategory) *Category {
	return &Category{
		repo,
	}
}

// Create insert information in the database
func (c Category) Create(ctx context.Context, cat domain.NewCategory) (*domain.Category, error) {
	return c.repo.Create(ctx, cat)
}

// Update  modifies data about a Category.
func (c Category) Update(ctx context.Context, cat domain.UpdateCategory) (*domain.Category, error) {
	return c.repo.Update(ctx, cat)
}

//Delete removes the category identified by a given ID.
func (c Category) Delete(ctx context.Context, categoryID int64) (err error) {
	return c.repo.Delete(ctx, categoryID)
}

// Query gets all Categories from the database.
func (c Category) Query(ctx context.Context) ([]domain.Category, error) {
	return c.repo.Query(ctx)
}

// QueryByID finds the category identified by a given ID.
func (c Category) QueryByID(ctx context.Context, categoryID int64) (*domain.Category, error) {
	return c.repo.QueryByID(ctx, categoryID)
}
