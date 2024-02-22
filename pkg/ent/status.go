// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/round"
	"github.com/scorify/backend/pkg/ent/status"
	"github.com/scorify/backend/pkg/ent/user"
)

// Status is the model entity for the Status schema.
type Status struct {
	config `json:"-"`
	// ID of the ent.
	// The uuid of a status
	ID uuid.UUID `json:"id"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// The error message of the status
	Error string `json:"error"`
	// The status of the status
	Status status.Status `json:"status"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusQuery when eager-loading is set.
	Edges        StatusEdges `json:"edges"`
	status_check *uuid.UUID
	status_round *uuid.UUID
	status_user  *uuid.UUID
	selectValues sql.SelectValues
}

// StatusEdges holds the relations/edges for other nodes in the graph.
type StatusEdges struct {
	// The check of a status
	Check *Check `json:"check"`
	// The round of a status
	Round *Round `json:"round"`
	// The user of a status
	User *User `json:"user"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CheckOrErr returns the Check value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StatusEdges) CheckOrErr() (*Check, error) {
	if e.loadedTypes[0] {
		if e.Check == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: check.Label}
		}
		return e.Check, nil
	}
	return nil, &NotLoadedError{edge: "check"}
}

// RoundOrErr returns the Round value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StatusEdges) RoundOrErr() (*Round, error) {
	if e.loadedTypes[1] {
		if e.Round == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: round.Label}
		}
		return e.Round, nil
	}
	return nil, &NotLoadedError{edge: "round"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StatusEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Status) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case status.FieldWeight:
			values[i] = new(sql.NullInt64)
		case status.FieldError, status.FieldStatus:
			values[i] = new(sql.NullString)
		case status.FieldCreateTime, status.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case status.FieldID:
			values[i] = new(uuid.UUID)
		case status.ForeignKeys[0]: // status_check
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case status.ForeignKeys[1]: // status_round
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case status.ForeignKeys[2]: // status_user
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Status fields.
func (s *Status) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case status.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case status.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case status.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case status.FieldError:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field error", values[i])
			} else if value.Valid {
				s.Error = value.String
			}
		case status.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = status.Status(value.String)
			}
		case status.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				s.Weight = int(value.Int64)
			}
		case status.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field status_check", values[i])
			} else if value.Valid {
				s.status_check = new(uuid.UUID)
				*s.status_check = *value.S.(*uuid.UUID)
			}
		case status.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field status_round", values[i])
			} else if value.Valid {
				s.status_round = new(uuid.UUID)
				*s.status_round = *value.S.(*uuid.UUID)
			}
		case status.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field status_user", values[i])
			} else if value.Valid {
				s.status_user = new(uuid.UUID)
				*s.status_user = *value.S.(*uuid.UUID)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Status.
// This includes values selected through modifiers, order, etc.
func (s *Status) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryCheck queries the "check" edge of the Status entity.
func (s *Status) QueryCheck() *CheckQuery {
	return NewStatusClient(s.config).QueryCheck(s)
}

// QueryRound queries the "round" edge of the Status entity.
func (s *Status) QueryRound() *RoundQuery {
	return NewStatusClient(s.config).QueryRound(s)
}

// QueryUser queries the "user" edge of the Status entity.
func (s *Status) QueryUser() *UserQuery {
	return NewStatusClient(s.config).QueryUser(s)
}

// Update returns a builder for updating this Status.
// Note that you need to call Status.Unwrap() before calling this method if this Status
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Status) Update() *StatusUpdateOne {
	return NewStatusClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Status entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Status) Unwrap() *Status {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Status is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Status) String() string {
	var builder strings.Builder
	builder.WriteString("Status(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("error=")
	builder.WriteString(s.Error)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", s.Weight))
	builder.WriteByte(')')
	return builder.String()
}

// StatusSlice is a parsable slice of Status.
type StatusSlice []*Status
