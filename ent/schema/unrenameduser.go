package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type UnrenamedUser struct {
	ent.Schema
}

// Fields of the UnrenamedUser.
func (UnrenamedUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Edges of the UnrenamedUser.
func (UnrenamedUser) Edges() []ent.Edge {
	return nil
}

func (UnrenamedUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
