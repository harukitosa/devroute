package schema

import (
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Unique().Immutable().NotEmpty(),
		field.String("email").Unique().NotEmpty().Validate(func(email string) error {
			_, err := mail.ParseAddress(email)
			return err
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
