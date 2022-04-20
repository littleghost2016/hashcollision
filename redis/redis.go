package redis

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/littleghost2016/hashcollision/config"
)

var (
	once sync.Once
	rdb  *redis.Client
)

func getNewClient() *redis.Client {
	once.Do(func() {
		rc := config.GetRedisConfig()
		rdb = redis.NewClient(&redis.Options{
			Addr:     rc.Address,
			Password: rc.Password, // no password set
			DB:       0,           // use default DB
		})
	})

	return rdb
}

func LPush[V any](listName string, value V) error {
	c := getNewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.LPush(ctx, listName, value).Result()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func RPush[V any](listName string, value V) error {
	c := getNewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := c.RPush(ctx, listName, value).Result()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func LPop(listName string) (string, error) {
	c := getNewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := c.LPop(ctx, listName).Result()
	if err != nil {
		return "", err
	} else {
		return value, nil
	}
}

func RPop(listName string) (string, error) {
	c := getNewClient()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	value, err := c.RPop(ctx, listName).Result()
	if err != nil {
		return "", err
	} else {
		return value, nil
	}
}
