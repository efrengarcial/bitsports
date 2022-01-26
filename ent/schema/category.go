package schema

import (
	"bitsports/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// CategoryMixin defines Fields
type CategoryMixin struct {
	entMixin.Schema
}

// Fields of the Category.
func (CategoryMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("code").
			NotEmpty().
			MaxLen(20).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(20)",      // Override Postgres.
			}),
		field.String("name").
			NotEmpty().
			MaxLen(100).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",      // Override Postgres.
			}),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return nil
}

// Mixin of the Category.
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CategoryMixin{},
		mixin.NewDatetime(),
	}
}
