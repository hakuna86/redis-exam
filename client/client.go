package client

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type Client struct {
	redis *redis.Client
}

func NewClient() *Client {
	c := new(Client)
	opt := &redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	c.redis = redis.NewClient(opt)
	return c
}

func (c *Client) Close() error {
	return c.redis.Close()
}

func (c *Client) Ping() error {
	return c.redis.Ping().Err()
}

func (c *Client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.redis.Set(key, value, expiration).Err()
}

func (c *Client) SetNotExist(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.redis.SetNX(key, value, expiration).Result()
}

func (c *Client) SetKeyMustExist(key string, value interface{}, expiration time.Duration) (bool, error) {
	return c.redis.SetXX(key, value, expiration).Result()
}

func (c *Client) Get(key string) (string, error) {
	return c.redis.Get(key).Result()
}
