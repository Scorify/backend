package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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
		field.Int("weight"),
	}
}

// Mixins of the Status.
func (Status) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
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
