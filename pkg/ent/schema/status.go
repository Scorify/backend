package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Status holds the schema definition for the Status entity.
type Status struct {
	ent.Schema
}

// Fields of the Status.
func (Status) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a status").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("error").
			StructTag(`json:"error"`).
			Comment("The error message of the status").
			Optional(),
		field.Enum("status").
			StructTag(`json:"status"`).
			Comment("The status of the status").
			Values("up", "down", "unknown").
			Default("unknown"),
		field.Time("updated_at").
			StructTag(`json:"updated_at"`).
			Comment("The update time of the status").
			Optional(),
		field.Int("weight"),
	}
}

// Edges of the Status.
func (Status) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("check", Check.Type).
			StructTag(`json:"check"`).
			Comment("The check of a status").
			Required().
			Unique(),
		edge.To("round", Round.Type).
			StructTag(`json:"round"`).
			Comment("The round of a status").
			Required().
			Unique(),
		edge.To("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user of a status").
			Required().
			Unique(),
	}
}
