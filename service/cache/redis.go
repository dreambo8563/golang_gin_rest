package cache

import (
	"strconv"

	"github.com/go-redis/redis"
	"vincent.com/golangginrest/config"
	"vincent.com/golangginrest/utils/logger"
)

var (
	client *redis.Client
	log    = logger.Log()
)

func init() {
	config := config.Item()
	add := config.Redis.Host + ":" + strconv.Itoa(config.Redis.Port)
	client = redis.NewClient(&redis.Options{
		Addr:     add,
		Password: config.Redis.Password,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf(err.Error(), "redis connect error")
	}
	log.Infoln("redis connected")
}

// Client will return the redis client instantce
func Client() *redis.Client {
	return client
}
