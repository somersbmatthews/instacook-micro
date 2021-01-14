package redis

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Client struct {
	redis.Client
}

func newClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func (c *Client) RedisGet(ctx context.Context, key string) (string, error) {
	val, redisErr := c.Get(ctx, "key").Result()
	err := returnProperRedisError(redisErr, val)
	if err != nil {
		return "", err
	}
	return val, nil
}

func (c *Client) RedisSet(ctx context.Context, key string, value string) error {
	val, redisErr := c.Set(ctx, key, value, 0).Result()
	err := returnProperRedisError(redisErr, val)
	if err != nil {
		return err
	}
	return nil
}

func returnProperRedisError(err error, val string) error {
	switch {
	case err == redis.Nil:
		return errors.New("key does not exist")
	case err != nil:
		errMsg := fmt.Sprintf("Get failed %v", err)
		return errors.New(errMsg)
	case val == "":
		return errors.New("value is empty")
	}
	return err
}
