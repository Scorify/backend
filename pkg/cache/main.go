package cache

import (
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
			Addr:     config.Redis.Url,
			Password: config.Redis.Password,
			DB:       0,
		}),
	}
}
