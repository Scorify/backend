// Code generated by ent, DO NOT EDIT.

package status

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldUpdateTime, v))
}

// Error applies equality check predicate on the "error" field. It's identical to ErrorEQ.
func Error(v string) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldError, v))
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v int) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldPoints, v))
}

// CheckID applies equality check predicate on the "check_id" field. It's identical to CheckIDEQ.
func CheckID(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldCheckID, v))
}

// RoundID applies equality check predicate on the "round_id" field. It's identical to RoundIDEQ.
func RoundID(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldRoundID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldUserID, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldUpdateTime, v))
}

// ErrorEQ applies the EQ predicate on the "error" field.
func ErrorEQ(v string) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldError, v))
}

// ErrorNEQ applies the NEQ predicate on the "error" field.
func ErrorNEQ(v string) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldError, v))
}

// ErrorIn applies the In predicate on the "error" field.
func ErrorIn(vs ...string) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldError, vs...))
}

// ErrorNotIn applies the NotIn predicate on the "error" field.
func ErrorNotIn(vs ...string) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldError, vs...))
}

// ErrorGT applies the GT predicate on the "error" field.
func ErrorGT(v string) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldError, v))
}

// ErrorGTE applies the GTE predicate on the "error" field.
func ErrorGTE(v string) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldError, v))
}

// ErrorLT applies the LT predicate on the "error" field.
func ErrorLT(v string) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldError, v))
}

// ErrorLTE applies the LTE predicate on the "error" field.
func ErrorLTE(v string) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldError, v))
}

// ErrorContains applies the Contains predicate on the "error" field.
func ErrorContains(v string) predicate.Status {
	return predicate.Status(sql.FieldContains(FieldError, v))
}

// ErrorHasPrefix applies the HasPrefix predicate on the "error" field.
func ErrorHasPrefix(v string) predicate.Status {
	return predicate.Status(sql.FieldHasPrefix(FieldError, v))
}

// ErrorHasSuffix applies the HasSuffix predicate on the "error" field.
func ErrorHasSuffix(v string) predicate.Status {
	return predicate.Status(sql.FieldHasSuffix(FieldError, v))
}

// ErrorIsNil applies the IsNil predicate on the "error" field.
func ErrorIsNil() predicate.Status {
	return predicate.Status(sql.FieldIsNull(FieldError))
}

// ErrorNotNil applies the NotNil predicate on the "error" field.
func ErrorNotNil() predicate.Status {
	return predicate.Status(sql.FieldNotNull(FieldError))
}

// ErrorEqualFold applies the EqualFold predicate on the "error" field.
func ErrorEqualFold(v string) predicate.Status {
	return predicate.Status(sql.FieldEqualFold(FieldError, v))
}

// ErrorContainsFold applies the ContainsFold predicate on the "error" field.
func ErrorContainsFold(v string) predicate.Status {
	return predicate.Status(sql.FieldContainsFold(FieldError, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldStatus, vs...))
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v int) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldPoints, v))
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v int) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldPoints, v))
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...int) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldPoints, vs...))
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...int) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldPoints, vs...))
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v int) predicate.Status {
	return predicate.Status(sql.FieldGT(FieldPoints, v))
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v int) predicate.Status {
	return predicate.Status(sql.FieldGTE(FieldPoints, v))
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v int) predicate.Status {
	return predicate.Status(sql.FieldLT(FieldPoints, v))
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v int) predicate.Status {
	return predicate.Status(sql.FieldLTE(FieldPoints, v))
}

// CheckIDEQ applies the EQ predicate on the "check_id" field.
func CheckIDEQ(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldCheckID, v))
}

// CheckIDNEQ applies the NEQ predicate on the "check_id" field.
func CheckIDNEQ(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldCheckID, v))
}

// CheckIDIn applies the In predicate on the "check_id" field.
func CheckIDIn(vs ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldCheckID, vs...))
}

// CheckIDNotIn applies the NotIn predicate on the "check_id" field.
func CheckIDNotIn(vs ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldCheckID, vs...))
}

// RoundIDEQ applies the EQ predicate on the "round_id" field.
func RoundIDEQ(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldRoundID, v))
}

// RoundIDNEQ applies the NEQ predicate on the "round_id" field.
func RoundIDNEQ(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldRoundID, v))
}

// RoundIDIn applies the In predicate on the "round_id" field.
func RoundIDIn(vs ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldRoundID, vs...))
}

// RoundIDNotIn applies the NotIn predicate on the "round_id" field.
func RoundIDNotIn(vs ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldRoundID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Status {
	return predicate.Status(sql.FieldNotIn(FieldUserID, vs...))
}

// HasCheck applies the HasEdge predicate on the "check" edge.
func HasCheck() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CheckTable, CheckColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCheckWith applies the HasEdge predicate on the "check" edge with a given conditions (other predicates).
func HasCheckWith(preds ...predicate.Check) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := newCheckStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConfig applies the HasEdge predicate on the "config" edge.
func HasConfig() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ConfigTable, ConfigColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConfigWith applies the HasEdge predicate on the "config" edge with a given conditions (other predicates).
func HasConfigWith(preds ...predicate.CheckConfig) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := newConfigStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRound applies the HasEdge predicate on the "round" edge.
func HasRound() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RoundTable, RoundColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoundWith applies the HasEdge predicate on the "round" edge with a given conditions (other predicates).
func HasRoundWith(preds ...predicate.Round) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := newRoundStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Status {
	return predicate.Status(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Status) predicate.Status {
	return predicate.Status(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Status) predicate.Status {
	return predicate.Status(sql.NotPredicates(p))
}
