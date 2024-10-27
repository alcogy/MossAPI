// Redis is Port number management Key-Value database.
// Key = service name / Value = port number
package redis

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


func FetchAllData() []KeyValue {
	db := connection()
	defer db.Close()

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

func CheckExistService(service string) bool {
	db := connection()
	defer db.Close()

	ctx := context.Background()

	port, err := db.Get(ctx, service).Result()
	if err != nil {
		return false
	}

	return port != ""
}

func GetPort(service string) (string, error) {
	db := connection()
	defer db.Close()

	ctx := context.Background()

	port, err := db.Get(ctx, service).Result()
	if err != nil {
		return "", err
	}

	return port, nil
}

func CheckPortNumberFree(port string) bool {
	kvs := FetchAllData()
	for _, v := range kvs {
		if v.Value == port {
			return false
		}
	}	
	return true
}

func SetService(service string, port string) error {
	db := connection()
	defer db.Close()

	ctx := context.Background()
	_, err := db.Set(ctx, service, port, 0).Result()

	return err
}

func DeleteService(service string) error {
	db := connection()
	defer db.Close()

	ctx := context.Background()
	_, err := db.Del(ctx, service).Result()
	
	return err
}