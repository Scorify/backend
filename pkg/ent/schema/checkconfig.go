package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CheckConfig holds the schema definition for the CheckConfig entity.
type CheckConfig struct {
	ent.Schema
}

// Fields of the CheckConfig.
func (CheckConfig) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a check configuration").
			Unique().
			Immutable().
			Default(uuid.New),
		field.JSON("config", map[string]interface{}{}).
			StructTag(`json:"config"`).
			Comment("The configuration of a check"),
	}
}

// Edges of the CheckConfig.
func (CheckConfig) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Check.Type).
			StructTag(`json:"check"`).
			Comment("The check this configuration belongs to").
			Required().
			Unique(),
		edge.To("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user this configuration belongs to").
			Required().
			Unique(),
	}
}
