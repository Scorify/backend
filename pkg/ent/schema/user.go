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
			NotEmpty(),
		field.String("password").
			Comment("The password hash of user password").
			Sensitive().
			NotEmpty(),
		field.Enum("role").
			StructTag(`json:"role"`).
			Comment("The role of the user").
			Values("admin", "user").
			Default("user").
			Immutable(),
		field.Int("number").
			StructTag(`json:"number"`).
			Comment("The number of the user").
			Optional().
			Unique().
			Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("configs", CheckConfig.Type).
			StructTag(`json:"config"`).
			Comment("The configuration of a check").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("user"),
		edge.From("status", Status.Type).
			StructTag(`json:"status"`).
			Comment("The status of a user").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("user"),
		edge.From("scorecaches", ScoreCache.Type).
			StructTag(`json:"scorecaches"`).
			Comment("The score caches of a user").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("user"),
	}
}
