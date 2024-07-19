package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	model "github.com/Udehlee/News-API-Wrapper/models"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// RedisConn returns a Redis client connection
func RedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// SetCache sets the news data in Redis cache
func SetCache(key string, newsData model.NewsData) error {
	client := RedisConn()

	newsJson, err := json.Marshal(&newsData)
	if err != nil {
		return fmt.Errorf("unable to marshal newsData: %w", err)
	}

	expirationTime := 1 * time.Hour

	if err := client.Set(ctx, key, newsJson, expirationTime).Err(); err != nil {
		return fmt.Errorf("unable to set cache: %w", err)
	}

	log.Printf("Cache set for key: %s", key)
	return nil
}

// GetCacheNews retrieves news data from Redis cache
func GetCacheNews(key string) (*model.NewsData, error) {
	client := RedisConn()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to get key: %w", err)
	}

	var newsData model.NewsData
	if err := json.Unmarshal([]byte(val), &newsData); err != nil {
		return nil, fmt.Errorf("unable to unmarshal newsData: %w", err)
	}

	log.Printf("Cache retrieved for key: %s", key)
	return &newsData, nil
}
