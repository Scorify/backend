package schema

import "entgo.io/ent"

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return nil
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return nil
}
