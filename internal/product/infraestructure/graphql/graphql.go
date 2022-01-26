package graphql

import (
	adapterGql "bitsports/internal/product/adapter/graphql"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewServer(s adapterGql.Schema) (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: s.Query(),
			Mutation:  s.Mutation(),
		},
	)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
		Playground: true,
	}), nil
}
