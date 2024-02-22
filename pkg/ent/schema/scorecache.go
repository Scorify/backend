package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ScoreCache holds the schema definition for the ScoreCache entity.
type ScoreCache struct {
	ent.Schema
}

// Fields of the ScoreCache.
func (ScoreCache) Fields() []ent.Field {
	return []ent.Field{
		field.Int("points").
			StructTag(`json:"points"`).
			Comment("The points of the round").
			NonNegative(),
	}
}

// Edges of the ScoreCache.
func (ScoreCache) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("round", Round.Type).
			StructTag(`json:"round"`).
			Comment("The round of a score cache").
			Required().
			Unique(),
		edge.To("user", User.Type).
			StructTag(`json:"user"`).
			Comment("The user of a score cache").
			Required().
			Unique(),
	}
}
