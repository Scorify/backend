// Code generated by ent, DO NOT EDIT.

package status

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the status type in the database.
	Label = "status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// FieldCheckID holds the string denoting the check_id field in the database.
	FieldCheckID = "check_id"
	// FieldRoundID holds the string denoting the round_id field in the database.
	FieldRoundID = "round_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeCheck holds the string denoting the check edge name in mutations.
	EdgeCheck = "check"
	// EdgeRound holds the string denoting the round edge name in mutations.
	EdgeRound = "round"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the status in the database.
	Table = "status"
	// CheckTable is the table that holds the check relation/edge.
	CheckTable = "status"
	// CheckInverseTable is the table name for the Check entity.
	// It exists in this package in order to avoid circular dependency with the "check" package.
	CheckInverseTable = "checks"
	// CheckColumn is the table column denoting the check relation/edge.
	CheckColumn = "check_id"
	// RoundTable is the table that holds the round relation/edge.
	RoundTable = "status"
	// RoundInverseTable is the table name for the Round entity.
	// It exists in this package in order to avoid circular dependency with the "round" package.
	RoundInverseTable = "rounds"
	// RoundColumn is the table column denoting the round relation/edge.
	RoundColumn = "round_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "status"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for status fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldError,
	FieldStatus,
	FieldPoints,
	FieldCheckID,
	FieldRoundID,
	FieldUserID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// PointsValidator is a validator for the "points" field. It is called by the builders before save.
	PointsValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusUnknown is the default value of the Status enum.
const DefaultStatus = StatusUnknown

// Status values.
const (
	StatusUp      Status = "up"
	StatusDown    Status = "down"
	StatusUnknown Status = "unknown"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusUp, StatusDown, StatusUnknown:
		return nil
	default:
		return fmt.Errorf("status: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Status queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByError orders the results by the error field.
func ByError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldError, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByPoints orders the results by the points field.
func ByPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoints, opts...).ToFunc()
}

// ByCheckID orders the results by the check_id field.
func ByCheckID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCheckID, opts...).ToFunc()
}

// ByRoundID orders the results by the round_id field.
func ByRoundID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoundID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByCheckField orders the results by check field.
func ByCheckField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCheckStep(), sql.OrderByField(field, opts...))
	}
}

// ByRoundField orders the results by round field.
func ByRoundField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoundStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newCheckStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CheckInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CheckTable, CheckColumn),
	)
}
func newRoundStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoundInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RoundTable, RoundColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
