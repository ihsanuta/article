package redis

import (
	"article/config"
	"fmt"
	"strconv"

	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func Connect() *redis.Client {
	redisHost := fmt.Sprintf("%s:%s", config.RedisConfig["host"], config.RedisConfig["port"])
	redisPass := fmt.Sprintf("%s", config.RedisConfig["password"])
	redisDB, _ := strconv.Atoi(fmt.Sprintf("%s", config.RedisConfig["database"]))

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: redisPass, // no password set
		DB:       redisDB,   // use default DB
	})

	greenOutput := color.New(color.FgGreen)
	successOutput := greenOutput.Add(color.Bold)

	successOutput.Println("")
	successOutput.Println("!!! Info")
	successOutput.Println(fmt.Sprintf("Successfully connected to redis %s", redisHost))
	successOutput.Println("")

	return rdb
}
