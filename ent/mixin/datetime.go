package mixin

import (
	"entgo.io/ent/dialect"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// NewDatetime creates a Mixin that includes create_at and updated_at
func NewDatetime() *DatetimeMixin {
	return &DatetimeMixin{}
}

// DatetimeMixin defines an ent Mixin
type DatetimeMixin struct {
	mixin.Schema
}

// Fields provides the created_at and updated_at field.
func (m DatetimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			SchemaType(map[string]string{
				dialect.Postgres: "timestamp",
			}).
			Immutable(),
	}
}
