package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// InjectSubmission holds the schema definition for the InjectSubmission entity.
type InjectSubmission struct {
	ent.Schema
}

// Fields of the InjectSubmission.
func (InjectSubmission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			StructTag(`json:"id"`).
			Comment("The uuid of an inject submission").
			Unique().
			Immutable().
			Default(uuid.New),
		field.Strings("files").
			StructTag(`json:"files"`).
			Comment("The files of the inject submission"),
	}
}

// Mixins of the InjectSubmission.
func (InjectSubmission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Edges of the InjectSubmission.
func (InjectSubmission) Edges() []ent.Edge {
	return []ent.Edge{}
}
