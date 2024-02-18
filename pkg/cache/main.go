package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/scorify/backend/pkg/config"
)

type Cache struct {
	Client *redis.Client
}

var (
	Client *Cache
)

func init() {
	Client = &Cache{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
			Password: config.Redis.Password,
			DB:       0,
		}),
	}
}
