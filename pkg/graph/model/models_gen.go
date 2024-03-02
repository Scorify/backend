// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type LoginOutput struct {
	Name     string `json:"name"`
	Token    string `json:"token"`
	Expires  int    `json:"expires"`
	Path     string `json:"path"`
	Domain   string `json:"domain"`
	Secure   bool   `json:"secure"`
	HTTPOnly bool   `json:"httpOnly"`
}

type Notification struct {
	Message string           `json:"message"`
	Type    NotificationType `json:"type"`
}

type Source struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type Subscription struct {
}

type EngineState string

const (
	EngineStateStopped EngineState = "stopped"
	EngineStateRunning EngineState = "running"
)

var AllEngineState = []EngineState{
	EngineStateStopped,
	EngineStateRunning,
}

func (e EngineState) IsValid() bool {
	switch e {
	case EngineStateStopped, EngineStateRunning:
		return true
	}
	return false
}

func (e EngineState) String() string {
	return string(e)
}

func (e *EngineState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EngineState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EngineState", str)
	}
	return nil
}

func (e EngineState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationType string

const (
	NotificationTypeDefault NotificationType = "default"
	NotificationTypeError   NotificationType = "error"
	NotificationTypeInfo    NotificationType = "info"
	NotificationTypeSuccess NotificationType = "success"
	NotificationTypeWarning NotificationType = "warning"
)

var AllNotificationType = []NotificationType{
	NotificationTypeDefault,
	NotificationTypeError,
	NotificationTypeInfo,
	NotificationTypeSuccess,
	NotificationTypeWarning,
}

func (e NotificationType) IsValid() bool {
	switch e {
	case NotificationTypeDefault, NotificationTypeError, NotificationTypeInfo, NotificationTypeSuccess, NotificationTypeWarning:
		return true
	}
	return false
}

func (e NotificationType) String() string {
	return string(e)
}

func (e *NotificationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationType", str)
	}
	return nil
}

func (e NotificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
