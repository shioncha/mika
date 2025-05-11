package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Tags_Posts holds the schema definition for the Tags_Posts entity.
type Tags_Posts struct {
	ent.Schema
}

// Fields of the Tags_Posts.
func (Tags_Posts) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("user_id"),
		field.Int("tag_id"),
		field.Int("post_id"),
	}
}

// Edges of the Tags_Posts.
func (Tags_Posts) Edges() []ent.Edge {
	return nil
}
