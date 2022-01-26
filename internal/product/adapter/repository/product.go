package repository

import (
	"bitsports/ent"
	entProduct "bitsports/ent/product"
	"bitsports/internal/product/domain"
	"context"
	"github.com/pkg/errors"
	"gopkg.in/jeevatkm/go-model.v1"
)

// product is the repository of domain.Product
type product struct {
	client *ent.Client
}

// NewProduct will create an object that represent the repository.IProduct interface
func NewProduct(client *ent.Client) *product {
	return &product{client}
}

func (r product) Create(ctx context.Context, prd domain.NewProduct) (*domain.Product, error) {
	p, err := r.client.
		Product.
		Create().
		SetName(prd.Name).
		SetCategoryID(prd.CategoryID).
		SetPrice(prd.Price).
		SetQuantity(prd.Quantity).
		Save(ctx)

	if err != nil {
		return nil, errors.WithStack(err)
	}
	return toProduct(p), nil
}

func (r product) Update(ctx context.Context, prd domain.UpdateProduct) (*domain.Product, error) {
	pu := r.client.Product.UpdateOneID(prd.ID)
	if prd.Name != nil {
		pu.SetName(*prd.Name)
	}
	if prd.Price != nil {
		pu.SetPrice(*prd.Price)
	}
	if prd.Quantity != nil {
		pu.SetQuantity(*prd.Quantity)
	}
	if prd.CategoryID != nil {
		pu.SetCategoryID(*prd.CategoryID)
	}

	p , err := pu.Save(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil,domain.ErrProductItemNotFound
		}

		return  nil,errors.WithStack(err)
	}

	return toProduct(p), nil
}

func (r product) Delete(ctx context.Context, productID int64) (err error) {
	err = r.client.Product.
		DeleteOneID(productID).
		Exec(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return domain.ErrProductItemNotFound
		}

		return  errors.WithStack(err)
	}
	return
}

func (r product) Query(ctx context.Context) ([]domain.Product, error) {
	products, err := r.client.
		Product.
		Query().
		WithCategory().
		All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return toProductSlice(products), nil
}

func (r product) QueryByID(ctx context.Context, productID int64) (*domain.Product, error) {
	q := r.client.Product.
		Query().
		WithCategory().
		Where(entProduct.IDEQ(productID))

	p, err := q.Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrProductItemNotFound
		}

		return nil, errors.WithStack(err)
	}

	return toProduct(p), nil
}

func toProduct(dbPrd *ent.Product) *domain.Product {
	var (
		p domain.Product
		c domain.Category
	)
	_ = model.Copy(&p, dbPrd)
	_ = model.Copy(&c, dbPrd.Edges.Category)
	p.Category = c

	return &p
}


func toProductSlice(dbPrds []*ent.Product) []domain.Product {
	prds := make([]domain.Product, len(dbPrds))
	for i, dbPrd := range dbPrds {
		prds[i] = *toProduct(dbPrd)
	}
	return prds
}
