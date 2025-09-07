package middlewares

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/labstack/echo/v4"
)

var mc = memcache.New("127.0.0.1:11211")

// CacheMiddleware caches only successful responses
func CacheMiddleware(key string, expiration time.Duration, fetch func(ctx context.Context) (interface{}, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Try getting data from cache
		if item, err := mc.Get(key); err == nil {
			var data interface{}
			if err := json.Unmarshal(item.Value, &data); err == nil {
				return Success(c, "Cache hit", data)
			}
		}

		// Fetch fresh data if cache miss
		data, err := fetch(c.Request().Context())
		if err != nil {
			return Error(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
		}

		// Cache only successful response
		bytes, _ := json.Marshal(data)
		mc.Set(&memcache.Item{
			Key:        key,
			Value:      bytes,
			Expiration: int32(expiration.Seconds()),
		})

		return Success(c, "Fetched successfully", data)
	}
}
