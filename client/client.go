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

func (c *Client) SetMap(key string, value map[string]interface{}) error {
	return c.redis.HMSet(key, value).Err()
}

func (c *Client) GetMap(key string) (map[string]string, error) {
	return c.redis.HGetAll(key).Result()
}

func (c *Client) SetSlice(key string, value ...interface{}) error {
	return c.redis.LPush(key, value...).Err()
}

func (c *Client) GetSlice(key string) ([]string, error) {
	return c.redis.LRange(key, 0, -1).Result()
}

func (c *Client) Get(key string) (string, error) {
	return c.redis.Get(key).Result()
}

func (c *Client) Publish(channel string, message interface{}) error {
	return c.redis.Publish(channel, message).Err()
}

func (c *Client) Subscribe(channel ...string) <-chan *redis.Message {
	return c.redis.Subscribe(channel...).Channel()
}
