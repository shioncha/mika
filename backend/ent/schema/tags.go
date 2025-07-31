package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

// Tags holds the schema definition for the Tags entity.
type Tags struct {
	ent.Schema
}

// Fields of the Tags.
func (Tags) Fields() []ent.Field {
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
	return []ent.Edge{
		edge.From("posts", Posts.Type).Ref("tags"),
	}
}
