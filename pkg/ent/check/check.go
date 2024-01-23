// Code generated by ent, DO NOT EDIT.

package check

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the check type in the database.
	Label = "check"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// EdgeConfig holds the string denoting the config edge name in mutations.
	EdgeConfig = "config"
	// Table holds the table name of the check in the database.
	Table = "checks"
	// ConfigTable is the table that holds the config relation/edge.
	ConfigTable = "check_configs"
	// ConfigInverseTable is the table name for the CheckConfig entity.
	// It exists in this package in order to avoid circular dependency with the "checkconfig" package.
	ConfigInverseTable = "check_configs"
	// ConfigColumn is the table column denoting the config relation/edge.
	ConfigColumn = "check_config_check"
)

// Columns holds all SQL columns for check fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSource,
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SourceValidator is a validator for the "source" field. It is called by the builders before save.
	SourceValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Check queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByConfigCount orders the results by config count.
func ByConfigCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newConfigStep(), opts...)
	}
}

// ByConfig orders the results by config terms.
func ByConfig(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newConfigStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newConfigStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ConfigInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ConfigTable, ConfigColumn),
	)
}
