// Code generated by ent, DO NOT EDIT.

package checkconfig

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldUpdateTime, v))
}

// CheckID applies equality check predicate on the "check_id" field. It's identical to CheckIDEQ.
func CheckID(v uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldCheckID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldUserID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldLTE(FieldUpdateTime, v))
}

// CheckIDEQ applies the EQ predicate on the "check_id" field.
func CheckIDEQ(v uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldCheckID, v))
}

// CheckIDNEQ applies the NEQ predicate on the "check_id" field.
func CheckIDNEQ(v uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNEQ(FieldCheckID, v))
}

// CheckIDIn applies the In predicate on the "check_id" field.
func CheckIDIn(vs ...uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldIn(FieldCheckID, vs...))
}

// CheckIDNotIn applies the NotIn predicate on the "check_id" field.
func CheckIDNotIn(vs ...uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNotIn(FieldCheckID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.CheckConfig {
	return predicate.CheckConfig(sql.FieldNotIn(FieldUserID, vs...))
}

// HasCheck applies the HasEdge predicate on the "check" edge.
func HasCheck() predicate.CheckConfig {
	return predicate.CheckConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CheckTable, CheckColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCheckWith applies the HasEdge predicate on the "check" edge with a given conditions (other predicates).
func HasCheckWith(preds ...predicate.Check) predicate.CheckConfig {
	return predicate.CheckConfig(func(s *sql.Selector) {
		step := newCheckStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.CheckConfig {
	return predicate.CheckConfig(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.CheckConfig {
	return predicate.CheckConfig(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CheckConfig) predicate.CheckConfig {
	return predicate.CheckConfig(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CheckConfig) predicate.CheckConfig {
	return predicate.CheckConfig(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CheckConfig) predicate.CheckConfig {
	return predicate.CheckConfig(sql.NotPredicates(p))
}
