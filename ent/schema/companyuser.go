package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// CompanyUser holds the schema definition for the CompanyUser entity.
type CompanyUser struct {
	ent.Schema
}

// Fields of the CompanyUser.
func (CompanyUser) Fields() []ent.Field {
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

// Edges of the CompanyUser.
func (CompanyUser) Edges() []ent.Edge {
	return nil
}
