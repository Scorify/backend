package graph

import (
	"github.com/scorify/backend/pkg/cache"
	"github.com/scorify/backend/pkg/ent"
	"github.com/scorify/backend/pkg/structs"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	Ent     *ent.Client
	Redis   *cache.Cache
	Checks  chan *ent.Status
	Minions *structs.Counter
}
