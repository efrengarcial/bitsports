package schema

import (
	"bitsports/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// ProductMixin defines Fields
type ProductMixin struct {
	entMixin.Schema
}

// Fields of the Product.
func (ProductMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("category_id"),
		field.String("name").
			NotEmpty().
			MaxLen(100).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",      // Override Postgres.
			}),
		field.Float("price").
			SchemaType(map[string]string{
				dialect.Postgres: "numeric(9,2)",      // Override Postgres.
			}),
		field.Int64("quantity").
			Positive(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type).
			Field("category_id").
			Unique().
			Required(),

	}
}

// Mixin of the Product.
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ProductMixin{},
		mixin.NewDatetime(),
	}
}
