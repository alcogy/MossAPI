package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type KeyValue struct {
	Key string
	Value string
}

func connection() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	return redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		Protocol: 2,
	})
}

func GetAllData() []KeyValue {
	db := connection()
	ctx := context.Background()

	keys, err := db.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}

	var kvs []KeyValue;
	for _, key := range keys {
		val, _ := db.Get(ctx, key).Result()
		kvs = append(kvs, KeyValue{key, val})
	}

	return kvs
}