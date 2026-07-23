package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/satudata/backend/pkg/cache"
)

// CacheRepository adalah wrapper untuk Redis cache dengan serialization JSON.
type CacheRepository struct {
	client *cache.RedisCache
	prefix string
}

// NewCacheRepository membuat instance baru CacheRepository.
func NewCacheRepository(client *cache.RedisCache) *CacheRepository {
	if client == nil {
		return nil
	}
	return &CacheRepository{
		client: client,
		prefix: "satudata:",
	}
}

// key mengembalikan key dengan prefix.
func (c *CacheRepository) key(k string) string {
	return c.prefix + k
}

// Get mengambil data dari cache dan mendecode ke target.
func (c *CacheRepository) Get(ctx context.Context, key string, target interface{}) error {
	if c.client == nil {
		return fmt.Errorf("cache not available")
	}

	data, err := c.client.Get(ctx, c.key(key))
	if err != nil {
		return fmt.Errorf("cache get error: %w", err)
	}
	if data == "" {
		return fmt.Errorf("cache miss")
	}

	if err := json.Unmarshal([]byte(data), target); err != nil {
		return fmt.Errorf("cache unmarshal error: %w", err)
	}

	return nil
}

// Set menyimpan data ke cache dengan TTL.
func (c *CacheRepository) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if c.client == nil {
		return fmt.Errorf("cache not available")
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("cache marshal error: %w", err)
	}

	return c.client.Set(ctx, c.key(key), string(data), ttl)
}

// Delete menghapus satu atau lebih keys dari cache.
func (c *CacheRepository) Delete(ctx context.Context, keys ...string) error {
	if c.client == nil {
		return fmt.Errorf("cache not available")
	}

	prefixedKeys := make([]string, len(keys))
	for i, k := range keys {
		prefixedKeys[i] = c.key(k)
	}

	return c.client.Delete(ctx, prefixedKeys...)
}

// Incr meng-increment nilai integer di cache.
func (c *CacheRepository) Incr(ctx context.Context, key string) (int64, error) {
	if c.client == nil {
		return 0, fmt.Errorf("cache not available")
	}

	return c.client.Incr(ctx, c.key(key))
}
