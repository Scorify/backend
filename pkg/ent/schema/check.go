package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Check holds the schema definition for the Check entity.
type Check struct {
	ent.Schema
}

// Fields of the Check.
func (Check) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of a check").
			Unique().
			Immutable().
			Default(uuid.New),
		field.String("name").
			StructTag(`json:"name"`).
			Comment("The name of the check").
			Unique().
			NotEmpty(),
		field.String("source").
			StructTag(`json:"source"`).
			Comment("The source of the check").
			NotEmpty(),
		field.JSON("config", map[string]interface{}{}).
			StructTag(`json:"config"`).
			Comment("The default configuration of a check"),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("config", CheckConfig.Type).
			StructTag(`json:"config"`).
			Comment("The configuration of a check").
			Ref("check"),
	}
}
