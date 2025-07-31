package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

// Posts holds the schema definition for the Posts entity.
type Posts struct {
	ent.Schema
}

// Fields of the Posts.
func (Posts) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(26).
			DefaultFunc(func() string {
				return ulid.Make().String()
			}).
			Immutable().
			Unique(),
		field.String("user_id").
			MaxLen(26).
			NotEmpty(),
		field.String("content").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Posts.
func (Posts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", Tags.Type),
	}
}
