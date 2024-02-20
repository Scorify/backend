// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/google/uuid"
	"github.com/scorify/backend/pkg/ent/check"
	"github.com/scorify/backend/pkg/ent/checkconfig"
	"github.com/scorify/backend/pkg/ent/schema"
	"github.com/scorify/backend/pkg/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	checkFields := schema.Check{}.Fields()
	_ = checkFields
	// checkDescName is the schema descriptor for name field.
	checkDescName := checkFields[1].Descriptor()
	// check.NameValidator is a validator for the "name" field. It is called by the builders before save.
	check.NameValidator = checkDescName.Validators[0].(func(string) error)
	// checkDescSource is the schema descriptor for source field.
	checkDescSource := checkFields[2].Descriptor()
	// check.SourceValidator is a validator for the "source" field. It is called by the builders before save.
	check.SourceValidator = checkDescSource.Validators[0].(func(string) error)
	// checkDescWeight is the schema descriptor for weight field.
	checkDescWeight := checkFields[3].Descriptor()
	// check.WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	check.WeightValidator = checkDescWeight.Validators[0].(func(int) error)
	// checkDescID is the schema descriptor for id field.
	checkDescID := checkFields[0].Descriptor()
	// check.DefaultID holds the default value on creation for the id field.
	check.DefaultID = checkDescID.Default.(func() uuid.UUID)
	checkconfigFields := schema.CheckConfig{}.Fields()
	_ = checkconfigFields
	// checkconfigDescID is the schema descriptor for id field.
	checkconfigDescID := checkconfigFields[0].Descriptor()
	// checkconfig.DefaultID holds the default value on creation for the id field.
	checkconfig.DefaultID = checkconfigDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescNumber is the schema descriptor for number field.
	userDescNumber := userFields[4].Descriptor()
	// user.NumberValidator is a validator for the "number" field. It is called by the builders before save.
	user.NumberValidator = userDescNumber.Validators[0].(func(int) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
