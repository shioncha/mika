package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("ulid").
			NotEmpty().
			Unique(),
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("password_hash").
			NotEmpty().
			Sensitive(),
		field.String("name").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("email_verify_at").
			Nillable().
			Optional(),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
