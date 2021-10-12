package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Recruitment holds the schema definition for the Recruitment entity.
type Recruitment struct {
	ent.Schema
}

// Fields of the Recruitment.
func (Recruitment) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("content").NotEmpty(),
	}
}

// Edges of the Recruitment.
func (Recruitment) Edges() []ent.Edge {
	return nil
}
