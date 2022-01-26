package graphql

import "github.com/graphql-go/graphql"

func categoryQuery(s Schema) graphql.Fields {
	return graphql.Fields{
		"category": &graphql.Field{
			Type:        categoryType,
			Description: "Get category by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: s.categoryResolver.QueryByID,
		},
		"categories": &graphql.Field{
			Type:        graphql.NewList(categoryType),
			Description: "Get category list",
			Resolve:     s.categoryResolver.Query,
		},
	}
}

func categoryMutation(s Schema) graphql.Fields {
	return graphql.Fields{
		"createCategory": &graphql.Field{
			Type:        categoryType,
			Description: "Create new category",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"code": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
			},
			Resolve: s.categoryResolver.Create,
		},
		"updateCategory": &graphql.Field{
			Type:        categoryType,
			Description: "Update category by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"code": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: s.categoryResolver.Update,
		},
		"deleteCategory": &graphql.Field{
			Type:        productType,
			Description: "Delete category by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: s.categoryResolver.Delete,
		},
	}
}
