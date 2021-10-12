package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Achievement holds the schema definition for the Achievement entity.
type Achievement struct {
	ent.Schema
}

// Fields of the Achievement.
func (Achievement) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Time("start_time"),
		field.Time("end_time"),
		field.String("content"),
	}
}

// Edges of the Achievement.
func (Achievement) Edges() []ent.Edge {
	return nil
}
