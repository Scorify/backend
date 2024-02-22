package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/structs"
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
		field.Int("weight").
			StructTag(`json:"weight"`).
			Comment("The weight of the check").
			NonNegative(),
		field.JSON("default_config", structs.CheckConfiguration{}).
			StructTag(`json:"default_config"`).
			Comment("The default configuration of a check"),
	}
}

// Mixins of the Check.
func (Check) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("configs", CheckConfig.Type).
			StructTag(`json:"config"`).
			Comment("The configuration of a check").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("check"),
		edge.From("statuses", Status.Type).
			StructTag(`json:"statuses"`).
			Comment("The statuses of a check").
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			).
			Ref("check"),
	}
}
