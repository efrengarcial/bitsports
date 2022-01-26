package graphql

import (
	"github.com/graphql-go/graphql"
)

var categoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"code": &graphql.Field{
				Type: graphql.String,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"quantity": &graphql.Field{
				Type: graphql.Int,
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
			},
			"category": &graphql.Field{
				Type: categoryType,
			},
		},
	},
)

// Schema is struct which has method for Query and Mutation. Please init this struct using constructor function.
type Schema struct {
	productResolver  Resolver
	categoryResolver Resolver
}

// Query initializes config schema query for graphql server.
func (s Schema) Query() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name:   "Query",
			Fields: merge(productQuery(s), categoryQuery(s)),
		})
}

// Mutation initializes config schema mutation for graphql server.
func (s Schema) Mutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Mutation",
		Fields: merge(productMutation(s), categoryMutation(s)),
	})
}

// NewSchema initializes Schema struct which takes productResolver as the argument.
func NewSchema(productResolver, categoryResolver Resolver) Schema {
	return Schema{productResolver, categoryResolver}
}

func merge(ms ...graphql.Fields) graphql.Fields {
	res := graphql.Fields{}
	for _, m := range ms {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
