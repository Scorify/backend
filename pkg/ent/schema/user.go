package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of the user").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("username").
			StructTag(`json:"username"`).
			Comment("The username of the user").
			Unique().
			Immutable().
			NotEmpty(),
		field.String("password").
			Comment("The password hash of user password").
			Sensitive().
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
