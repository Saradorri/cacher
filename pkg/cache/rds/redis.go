package rds

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type CacheClient interface {
	Get(key string) interface{}
	Set(key string, value interface{}, ttl time.Duration)
	Delete(key string)
}

type cacheClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewClient(addr string) CacheClient {
	rs := &cacheClient{ctx: context.Background()}

	if rs.client != nil {
		if _, e := rs.client.Ping(rs.ctx).Result(); e == nil {
			return rs
		}
	}
	rs.client = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	return rs
}

func (rdc *cacheClient) Set(key string, value interface{}, ttl time.Duration) {
	v, _ := json.Marshal(value)
	err := rdc.client.Set(rdc.ctx, key, v, ttl).Err()

	if err != nil {
		log.Printf("redis set data error: %s\n", err.Error())
		time.Sleep(3 * time.Second)
		rdc.client.Set(rdc.ctx, key, v, ttl)
	}
}

func (rdc *cacheClient) Get(key string) interface{} {
	r, err := rdc.client.Get(rdc.ctx, key).Bytes()
	if err != nil {
		log.Printf("redis get data error: %s\n", err.Error())
	}

	var value string
	if !(errors.Is(err, redis.Nil) || err != nil) {
		_ = json.Unmarshal(r, &value)
		return value
	}
	return nil
}

func (rdc *cacheClient) Delete(key string) {
	rdc.client.Del(rdc.ctx, key)
}
