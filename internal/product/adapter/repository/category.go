package repository

import (
	"bitsports/ent"
	entCategory "bitsports/ent/category"
	"bitsports/internal/product/domain"
	"context"
	"github.com/pkg/errors"
	"gopkg.in/jeevatkm/go-model.v1"
)

// category is the repository of domain.Category
type category struct {
	client *ent.Client
}

// NewCategory will create an object that represent the repository.ICategory interface
func NewCategory(client *ent.Client) *category {
	return &category{client}
}

func (c category) Create(ctx context.Context, cat domain.NewCategory) (*domain.Category, error) {
	ca, err := c.client.
		Category.
		Create().
		SetName(cat.Name).
		SetCode(cat.Code).
		Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return toCategory(ca), nil
}

func (c category) Update(ctx context.Context, cat domain.UpdateCategory) (*domain.Category, error) {
	cu := c.client.Category.UpdateOneID(cat.ID)
	if cat.Name != nil {
		cu.SetName(*cat.Name)
	}
	if cat.Code != nil {
		cu.SetCode(*cat.Code)
	}

	ca , err := cu.Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil,domain.ErrCategoryItemNotFound
		}

		return  nil,errors.WithStack(err)
	}

	return toCategory(ca), nil
}

func (c category) Delete(ctx context.Context, categoryID int64) (err error) {
	err = c.client.Category.
		DeleteOneID(categoryID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return domain.ErrCategoryItemNotFound
		}

		return  errors.WithStack(err)
	}
	return
}

func (c category) Query(ctx context.Context) ([]domain.Category, error) {
	categories, err := c.client.
		Category.
		Query().
		All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return toCategorySlice(categories), nil
}

func (c category) QueryByID(ctx context.Context, categoryID int64) (*domain.Category, error) {
	q := c.client.Category.
		Query().
		Where(entCategory.IDEQ(categoryID))

	ca, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrCategoryItemNotFound
		}

		return nil, errors.WithStack(err)
	}

	return toCategory(ca), nil
}

func toCategory(dbCat *ent.Category) *domain.Category {
	var c domain.Category
	_ = model.Copy(&c, dbCat)
	return &c
}

func toCategorySlice(dbCats []*ent.Category) []domain.Category {
	categories := make([]domain.Category, len(dbCats))
	for i, dbCat := range dbCats {
		categories[i] = *toCategory(dbCat)
	}
	return categories
}
