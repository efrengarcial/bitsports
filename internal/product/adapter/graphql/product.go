package graphql

import "github.com/graphql-go/graphql"

func productQuery(s Schema) graphql.Fields {
	return graphql.Fields{
		/* Get (read) single product by id
		   http://localhost:8080/graphql?query={product(id:1){name,price}}
		*/
		"product": &graphql.Field{
			Type:        productType,
			Description: "Get product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: s.productResolver.QueryByID,
		},
		/* Get (read) product list
		   http://localhost:8080/graphql?query={products{id,name,price}}
		*/
		"products": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Get product list",
			Resolve:     s.productResolver.Query,
		},
	}
}

func productMutation(s Schema) graphql.Fields {
	return graphql.Fields{
		/* Create new product item
		http://localhost:8080/graphql?query=mutation+_{createProduct(name:"xxx" ,price:1.99){id,name,info,price}}
		*/
		"createProduct": &graphql.Field{
			Type:        productType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"quantity": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"categoryId": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: s.productResolver.Create,
		},

		/* Update product by id
		   http://localhost:8080/graphql?query=mutation+_{updateProduct(id:1,price:3.95){id,name,info,price}}
		*/
		"updateProduct": &graphql.Field{
			Type:        productType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"price": &graphql.ArgumentConfig{
					Type: graphql.Float,
				},
				"quantity": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"categoryId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: s.productResolver.Update,
		},

		/* Delete product by id
		   http://localhost:8080/graphql?query=mutation+_{delete(id:1){id,name,price}}
		*/
		"deleteProduct": &graphql.Field{
			Type:        productType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: s.productResolver.Delete,
		},
	}
}
