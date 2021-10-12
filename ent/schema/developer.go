package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Developer holds the schema definition for the Developer entity.
type Developer struct {
	ent.Schema
}

// Fields of the Developer.
func (Developer) Fields() []ent.Field {
	return []ent.Field{
		field.String("last_name").NotEmpty(),
		field.String("first_name").NotEmpty(),
		field.String("last_name_furigana").NotEmpty(),
		field.String("first_name_furigana").NotEmpty(),
		field.String("profile_name").NotEmpty(),
		field.String("icon_url").NotEmpty(),
		field.Enum("gender").Values("male", "female", "other"),
	}
}

// Edges of the Developer.
func (Developer) Edges() []ent.Edge {
	return nil
}
