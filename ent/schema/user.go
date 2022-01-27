package schema

import (
	"bitsports/ent/mixin"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	entMixin "entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// UserMixin defines Fields
type UserMixin struct {
	entMixin.Schema
}


// Fields of the User.
func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").
			NotEmpty().
			MaxLen(100).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",      // Override Postgres.
			}),
		field.String("email").
			NotEmpty().
			MaxLen(100).
			SchemaType(map[string]string{
				dialect.Postgres: "varchar(100)",      // Override Postgres.
			}),
		field.Bytes("password_hash").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}


// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{},
		mixin.NewDatetime(),
	}
}
