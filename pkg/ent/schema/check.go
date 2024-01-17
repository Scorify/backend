package schema

import "entgo.io/ent"

// Check holds the schema definition for the Check entity.
type Check struct {
	ent.Schema
}

// Fields of the Check.
func (Check) Fields() []ent.Field {
	return nil
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return nil
}
