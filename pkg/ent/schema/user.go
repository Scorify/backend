package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
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
		field.Enum("role").
			StructTag(`json:"role"`).
			Comment("The role of the user").
			Values("admin", "user").
			Default("user"),
		field.Int("number").
			StructTag(`json:"number"`).
			Comment("The number of the user").
			Immutable().
			Optional().
			Unique().
			Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("configs", CheckConfig.Type).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			StructTag(`json:"config"`).
			Comment("The configuration of a check").
			Ref("user"),
	}
}
