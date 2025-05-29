package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dotenv-org/godotenvvault"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient() *redis.Client {

	err := godotenvvault.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}

	fmt.Println("Connecting to redis server on:", os.Getenv("REDIS_HOST"))

	redisDB := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
		Username: "default",
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})
	return redisDB
}

func SetKey(context *context.Context, redisDB *redis.Client, key string, value string, ttl int) {
	fmt.Println("Setting key", key, "to", value, "in Redis")

	redisDB.Set(*context, key, value, 0)
	fmt.Println("The key", key, "has been set to", value, "successfully")

}

func GetLongURL(context *context.Context, redisDB *redis.Client, shortURL string) (string, error) {
	longURL, err := redisDB.Get(*context, shortURL).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("short URL not found")
	} else if err != nil {
		return "", fmt.Errorf("failed to retrieve from Redis: %v", err)
	}

	return longURL, nil
}
