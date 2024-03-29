// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/scorify/backend/pkg/ent/status"
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

type RoundUpdateScoreboard struct {
	Round    int  `json:"round"`
	Complete bool `json:"complete"`
}

type ScoreUpdateScoreboard struct {
	Team   int `json:"team"`
	Round  int `json:"round"`
	Points int `json:"points"`
}

type ScoreboardUpdate struct {
	StatusUpdate []*StatusUpdateScoreboard `json:"statusUpdate,omitempty"`
	RoundUpdate  []*RoundUpdateScoreboard  `json:"roundUpdate,omitempty"`
	ScoreUpdate  []*ScoreUpdateScoreboard  `json:"scoreUpdate,omitempty"`
}

type Source struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type StatusUpdateScoreboard struct {
	Team   int           `json:"team"`
	Round  int           `json:"round"`
	Check  string        `json:"check"`
	Status status.Status `json:"status"`
}

type Subscription struct {
}

type EngineState string

const (
	EngineStatePaused   EngineState = "paused"
	EngineStateWaiting  EngineState = "waiting"
	EngineStateRunning  EngineState = "running"
	EngineStateStopping EngineState = "stopping"
)

var AllEngineState = []EngineState{
	EngineStatePaused,
	EngineStateWaiting,
	EngineStateRunning,
	EngineStateStopping,
}

func (e EngineState) IsValid() bool {
	switch e {
	case EngineStatePaused, EngineStateWaiting, EngineStateRunning, EngineStateStopping:
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
