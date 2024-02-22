// Code generated by ent, DO NOT EDIT.

package round

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the round type in the database.
	Label = "round"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldComplete holds the string denoting the complete field in the database.
	FieldComplete = "complete"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// EdgeStatuses holds the string denoting the statuses edge name in mutations.
	EdgeStatuses = "statuses"
	// EdgeScorecaches holds the string denoting the scorecaches edge name in mutations.
	EdgeScorecaches = "scorecaches"
	// Table holds the table name of the round in the database.
	Table = "rounds"
	// StatusesTable is the table that holds the statuses relation/edge.
	StatusesTable = "status"
	// StatusesInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusesInverseTable = "status"
	// StatusesColumn is the table column denoting the statuses relation/edge.
	StatusesColumn = "status_round"
	// ScorecachesTable is the table that holds the scorecaches relation/edge.
	ScorecachesTable = "score_caches"
	// ScorecachesInverseTable is the table name for the ScoreCache entity.
	// It exists in this package in order to avoid circular dependency with the "scorecache" package.
	ScorecachesInverseTable = "score_caches"
	// ScorecachesColumn is the table column denoting the scorecaches relation/edge.
	ScorecachesColumn = "score_cache_round"
)

// Columns holds all SQL columns for round fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldNumber,
	FieldComplete,
	FieldPoints,
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
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(int) error
	// DefaultComplete holds the default value on creation for the "complete" field.
	DefaultComplete bool
	// PointsValidator is a validator for the "points" field. It is called by the builders before save.
	PointsValidator func(int) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Round queries.
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

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByComplete orders the results by the complete field.
func ByComplete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComplete, opts...).ToFunc()
}

// ByPoints orders the results by the points field.
func ByPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoints, opts...).ToFunc()
}

// ByStatusesCount orders the results by statuses count.
func ByStatusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatusesStep(), opts...)
	}
}

// ByStatuses orders the results by statuses terms.
func ByStatuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByScorecachesCount orders the results by scorecaches count.
func ByScorecachesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newScorecachesStep(), opts...)
	}
}

// ByScorecaches orders the results by scorecaches terms.
func ByScorecaches(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScorecachesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newStatusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, StatusesTable, StatusesColumn),
	)
}
func newScorecachesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScorecachesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ScorecachesTable, ScorecachesColumn),
	)
}
