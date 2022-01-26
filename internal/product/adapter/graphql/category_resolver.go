package graphql

import (
	"bitsports/internal/product/application/usecase"
	"bitsports/internal/product/domain"
	"errors"
	"github.com/graphql-go/graphql"
)

type categoryResolver struct {
	uc *usecase.Category
}

func NewCategoryResolver(uc *usecase.Category) *categoryResolver {
	return &categoryResolver{uc}
}

func (r categoryResolver) Query(params graphql.ResolveParams) (interface{}, error) {
	results, err := r.uc.Query(params.Context)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (r categoryResolver) QueryByID(params graphql.ResolveParams) (interface{}, error) {
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

func (r categoryResolver) Update(params graphql.ResolveParams) (interface{}, error) {
	var (
		id       int
		ok       bool
		category *domain.Category
		err      error
	)

	if id, ok = params.Args["id"].(int); !ok || id == 0 {
		return nil, errors.New("id is not integer or zero")
	}

	name, nameOk := params.Args["name"].(string)
	code, codeOk := params.Args["code"].(string)

	updateCategory := domain.UpdateCategory{ID: int64(id)}

	if nameOk {
		updateCategory.Name = &name
	}

	if codeOk {
		updateCategory.Code = &code
	}

	if category, err = r.uc.Update(params.Context, updateCategory); err != nil {
		return nil, err
	}

	return *category, nil
}

func (r categoryResolver) Delete(params graphql.ResolveParams) (interface{}, error) {
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

func (r categoryResolver) Create(params graphql.ResolveParams) (interface{}, error) {
	var (
		name, code string
		ok         bool
		category   *domain.Category
		err        error
	)

	if name, ok = params.Args["name"].(string); !ok || name == "" {
		return nil, errors.New("name is empty or not string")
	}

	if code, ok = params.Args["code"].(string); !ok || name == "" {
		return nil, errors.New("code is empty or not string")
	}

	storedCategory := domain.NewCategory{
		Name: name,
		Code: code,
	}

	if category, err = r.uc.Create(params.Context, storedCategory); err != nil {
		return nil, err
	}

	return *category, nil
}
