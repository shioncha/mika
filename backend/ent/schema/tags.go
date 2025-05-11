package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tags holds the schema definition for the Tags entity.
type Tags struct {
	ent.Schema
}

// Fields of the Tags.
func (Tags) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("ulid").
			NotEmpty().
			Unique(),
		field.Int("user_id"),
		field.String("tag").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Tags.
func (Tags) Edges() []ent.Edge {
	return nil
}
