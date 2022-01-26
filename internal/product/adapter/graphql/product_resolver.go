package graphql

import (
	"bitsports/internal/product/application/usecase"
	"bitsports/internal/product/domain"
	"errors"
	"github.com/graphql-go/graphql"
)

type productResolver struct {
	uc *usecase.Product
}

func NewProductResolver(uc *usecase.Product) *productResolver {
	return &productResolver{uc}
}

func (r productResolver) Query(params graphql.ResolveParams) (interface{}, error) {
	results, err := r.uc.Query(params.Context)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r productResolver) QueryByID(params graphql.ResolveParams) (interface{}, error) {
	var (
		id int
		ok bool
	)

	if id, ok = params.Args["id"].(int); !ok || id == 0 {
		return nil, errors.New("id is not integer or zero")
	}

	result, err := r.uc.QueryByID(params.Context, int64(id))
	if err != nil {
		return nil, err
	}
	return *result, nil
}

func (r productResolver) Update(params graphql.ResolveParams) (interface{}, error) {
	var (
		id      int
		ok      bool
		product *domain.Product
		err     error
	)

	if id, ok = params.Args["id"].(int); !ok || id == 0 {
		return nil, errors.New("id is not integer or zero")
	}

	name, nameOk := params.Args["name"].(string)
	price, priceOk := params.Args["price"].(float64)
	quantity, quantityOk := params.Args["quantity"].(int)
	categoryID, categoryIdOk := params.Args["categoryId"].(int)

	updatedProduct := domain.UpdateProduct{ID: int64(id)}

	if nameOk {
		updatedProduct.Name = &name
	}

	if priceOk {
		updatedProduct.Price = &price
	}

	if quantityOk {
		q := int64(quantity)
		updatedProduct.Quantity = &q
	}

	if categoryIdOk {
		c := int64(categoryID)
		updatedProduct.Quantity = &c
	}

	if product, err = r.uc.Update(params.Context, updatedProduct); err != nil {
		return nil, err
	}

	return *product, nil
}

func (r productResolver) Delete(params graphql.ResolveParams) (interface{}, error) {
	var (
		id int
		ok bool
	)

	if id, ok = params.Args["id"].(int); !ok || id == 0 {
		return nil, errors.New("id is not integer or zero")
	}

	if err := r.uc.Delete(params.Context, int64(id)); err != nil {
		return nil, err
	}

	return id, nil
}

func (r productResolver) Create(params graphql.ResolveParams) (interface{}, error) {
	var (
		name                 string
		price                float64
		quantity, categoryID int
		ok                   bool
		product              *domain.Product
		err                  error
	)

	if name, ok = params.Args["name"].(string); !ok || name == "" {
		return nil, errors.New("name is empty or not string")
	}

	if price, ok = params.Args["price"].(float64); !ok {
		return nil, errors.New("price is not number")
	}

	if quantity, ok = params.Args["quantity"].(int); !ok {
		return nil, errors.New("quantity is not number")
	}

	if categoryID, ok = params.Args["categoryId"].(int); !ok {
		return nil, errors.New("categoryId is not number")
	}

	storedProduct := domain.NewProduct{
		Name:       name,
		Price:      price,
		Quantity:   int64(quantity),
		CategoryID: int64(categoryID),
	}

	if product, err = r.uc.Create(params.Context, storedProduct); err != nil {
		return nil, err
	}

	return *product, nil
}
