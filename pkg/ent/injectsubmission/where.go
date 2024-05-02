// Code generated by ent, DO NOT EDIT.

package injectsubmission

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldUpdateTime, v))
}

// InjectID applies equality check predicate on the "inject_id" field. It's identical to InjectIDEQ.
func InjectID(v uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldInjectID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldUserID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldLTE(FieldUpdateTime, v))
}

// InjectIDEQ applies the EQ predicate on the "inject_id" field.
func InjectIDEQ(v uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldInjectID, v))
}

// InjectIDNEQ applies the NEQ predicate on the "inject_id" field.
func InjectIDNEQ(v uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNEQ(FieldInjectID, v))
}

// InjectIDIn applies the In predicate on the "inject_id" field.
func InjectIDIn(vs ...uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldIn(FieldInjectID, vs...))
}

// InjectIDNotIn applies the NotIn predicate on the "inject_id" field.
func InjectIDNotIn(vs ...uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNotIn(FieldInjectID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.FieldNotIn(FieldUserID, vs...))
}

// HasInject applies the HasEdge predicate on the "inject" edge.
func HasInject() predicate.InjectSubmission {
	return predicate.InjectSubmission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InjectTable, InjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInjectWith applies the HasEdge predicate on the "inject" edge with a given conditions (other predicates).
func HasInjectWith(preds ...predicate.Inject) predicate.InjectSubmission {
	return predicate.InjectSubmission(func(s *sql.Selector) {
		step := newInjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.InjectSubmission {
	return predicate.InjectSubmission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.InjectSubmission {
	return predicate.InjectSubmission(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InjectSubmission) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InjectSubmission) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.InjectSubmission) predicate.InjectSubmission {
	return predicate.InjectSubmission(sql.NotPredicates(p))
}
