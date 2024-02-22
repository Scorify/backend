// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/structs"
)

// Check is the model entity for the Check schema.
type Check struct {
	config `json:"-"`
	// ID of the ent.
	// The uuid of a check
	ID uuid.UUID `json:"id"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The name of the check
	Name string `json:"name"`
	// The source of the check
	Source string `json:"source"`
	// The weight of the check
	Weight int `json:"weight"`
	// The default configuration of a check
	DefaultConfig structs.CheckConfiguration `json:"default_config"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CheckQuery when eager-loading is set.
	Edges        CheckEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CheckEdges holds the relations/edges for other nodes in the graph.
type CheckEdges struct {
	// The configuration of a check
	Configs []*CheckConfig `json:"config"`
	// The statuses of a check
	Statuses []*Status `json:"statuses"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ConfigsOrErr returns the Configs value or an error if the edge
// was not loaded in eager-loading.
func (e CheckEdges) ConfigsOrErr() ([]*CheckConfig, error) {
	if e.loadedTypes[0] {
		return e.Configs, nil
	}
	return nil, &NotLoadedError{edge: "configs"}
}

// StatusesOrErr returns the Statuses value or an error if the edge
// was not loaded in eager-loading.
func (e CheckEdges) StatusesOrErr() ([]*Status, error) {
	if e.loadedTypes[1] {
		return e.Statuses, nil
	}
	return nil, &NotLoadedError{edge: "statuses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Check) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case check.FieldDefaultConfig:
			values[i] = new([]byte)
		case check.FieldWeight:
			values[i] = new(sql.NullInt64)
		case check.FieldName, check.FieldSource:
			values[i] = new(sql.NullString)
		case check.FieldCreateTime, check.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case check.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Check fields.
func (c *Check) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case check.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case check.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case check.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case check.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case check.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				c.Source = value.String
			}
		case check.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				c.Weight = int(value.Int64)
			}
		case check.FieldDefaultConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field default_config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.DefaultConfig); err != nil {
					return fmt.Errorf("unmarshal field default_config: %w", err)
				}
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Check.
// This includes values selected through modifiers, order, etc.
func (c *Check) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryConfigs queries the "configs" edge of the Check entity.
func (c *Check) QueryConfigs() *CheckConfigQuery {
	return NewCheckClient(c.config).QueryConfigs(c)
}

// QueryStatuses queries the "statuses" edge of the Check entity.
func (c *Check) QueryStatuses() *StatusQuery {
	return NewCheckClient(c.config).QueryStatuses(c)
}

// Update returns a builder for updating this Check.
// Note that you need to call Check.Unwrap() before calling this method if this Check
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Check) Update() *CheckUpdateOne {
	return NewCheckClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Check entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Check) Unwrap() *Check {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Check is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Check) String() string {
	var builder strings.Builder
	builder.WriteString("Check(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(c.Source)
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", c.Weight))
	builder.WriteString(", ")
	builder.WriteString("default_config=")
	builder.WriteString(fmt.Sprintf("%v", c.DefaultConfig))
	builder.WriteByte(')')
	return builder.String()
}

// Checks is a parsable slice of Check.
type Checks []*Check
