package graphql

import "github.com/graphql-go/graphql"

type Resolver interface {
	Query(params graphql.ResolveParams) (interface{}, error)
	QueryByID(params graphql.ResolveParams) (interface{}, error)

	Update(params graphql.ResolveParams) (interface{}, error)
	Delete(params graphql.ResolveParams) (interface{}, error)
	Create(params graphql.ResolveParams) (interface{}, error)
}

