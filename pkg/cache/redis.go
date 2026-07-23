package cache

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/satudata/backend/internal/config"
)

var rdb *redis.Client

// CacheInterface mendefinisikan operasi caching dasar.
type CacheInterface interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, key string) (bool, error)
	Flush(ctx context.Context) error
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, ttl time.Duration) error
	Ping(ctx context.Context) error
}

// RedisCache adalah implementasi CacheInterface menggunakan Redis.
type RedisCache struct {
	client *redis.Client
}

// Init menginisialisasi koneksi Redis.
func Init(cfg config.RedisConfig) (*RedisCache, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)

	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	rdb = client

	// Test koneksi
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	cache := &RedisCache{client: client}
	log.Println("[Cache] Successfully connected to Redis")
	return cache, nil
}

// Get mengambil value dari cache berdasarkan key.
func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Key tidak ditemukan, bukan error
	}
	return val, err
}

// Set menyimpan value ke cache dengan TTL.
func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

// Delete menghapus satu atau lebih keys dari cache.
func (c *RedisCache) Delete(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// Exists mengecek apakah key ada di cache.
func (c *RedisCache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, key).Result()
	return n > 0, err
}

// Flush menghapus semua keys di database Redis saat ini.
func (c *RedisCache) Flush(ctx context.Context) error {
	return c.client.FlushDB(ctx).Err()
}

// Incr meng-increment nilai integer di cache.
func (c *RedisCache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

// Expire mengatur TTL untuk key yang sudah ada.
func (c *RedisCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return c.client.Expire(ctx, key, ttl).Err()
}

// Ping memeriksa koneksi Redis.
func (c *RedisCache) Ping(ctx context.Context) error {
	return c.client.Ping(ctx).Err()
}

// GetClient mengembalikan Redis client langsung (untuk advanced use cases).
func GetClient() *redis.Client {
	if rdb == nil {
		log.Fatal("[Cache] Redis not initialized. Call Init() first.")
	}
	return rdb
}

// Close menutup koneksi Redis.
func Close() error {
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}
