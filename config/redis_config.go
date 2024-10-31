package config

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client;

func ConnectToRedis() error {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:   db,
	})

	_, err = rdb.Ping(context.Background()).Result()

	if err != nil {
		return err
	}

	Redis = rdb
	log.Println("Connected to Redis!")

	return nil

}