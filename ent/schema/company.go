package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("location").NotEmpty(),
		field.String("postal_code").NotEmpty(),
		field.String("phone_number").NotEmpty(),
		field.String("company_id").NotEmpty(),
		field.String("introduction").NotEmpty(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}
