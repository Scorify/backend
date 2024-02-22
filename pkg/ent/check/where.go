// Code generated by ent, DO NOT EDIT.

package check

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Check {
	return predicate.Check(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldName, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldSource, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldWeight, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Check {
	return predicate.Check(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Check {
	return predicate.Check(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Check {
	return predicate.Check(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Check {
	return predicate.Check(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Check {
	return predicate.Check(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Check {
	return predicate.Check(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Check {
	return predicate.Check(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Check {
	return predicate.Check(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Check {
	return predicate.Check(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Check {
	return predicate.Check(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Check {
	return predicate.Check(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Check {
	return predicate.Check(sql.FieldContainsFold(FieldName, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Check {
	return predicate.Check(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Check {
	return predicate.Check(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Check {
	return predicate.Check(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Check {
	return predicate.Check(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Check {
	return predicate.Check(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Check {
	return predicate.Check(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Check {
	return predicate.Check(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Check {
	return predicate.Check(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Check {
	return predicate.Check(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Check {
	return predicate.Check(sql.FieldHasSuffix(FieldSource, v))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Check {
	return predicate.Check(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Check {
	return predicate.Check(sql.FieldContainsFold(FieldSource, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.Check {
	return predicate.Check(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.Check {
	return predicate.Check(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.Check {
	return predicate.Check(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.Check {
	return predicate.Check(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.Check {
	return predicate.Check(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.Check {
	return predicate.Check(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.Check {
	return predicate.Check(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.Check {
	return predicate.Check(sql.FieldLTE(FieldWeight, v))
}

// HasConfigs applies the HasEdge predicate on the "configs" edge.
func HasConfigs() predicate.Check {
	return predicate.Check(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ConfigsTable, ConfigsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigsWith applies the HasEdge predicate on the "configs" edge with a given conditions (other predicates).
func HasConfigsWith(preds ...predicate.CheckConfig) predicate.Check {
	return predicate.Check(func(s *sql.Selector) {
		step := newConfigsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasStatuses applies the HasEdge predicate on the "statuses" edge.
func HasStatuses() predicate.Check {
	return predicate.Check(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, StatusesTable, StatusesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStatusesWith applies the HasEdge predicate on the "statuses" edge with a given conditions (other predicates).
func HasStatusesWith(preds ...predicate.Status) predicate.Check {
	return predicate.Check(func(s *sql.Selector) {
		step := newStatusesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Check) predicate.Check {
	return predicate.Check(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Check) predicate.Check {
	return predicate.Check(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Check) predicate.Check {
	return predicate.Check(sql.NotPredicates(p))
}
