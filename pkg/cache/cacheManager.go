package cache

import (
	"cacher/pkg/cache/rds"
	"cacher/pkg/hashing"
	"fmt"
)

type CacheManager interface {
	GetRedisClient(key string) rds.CacheClient
}

type cacheManager struct {
	hashRing *hashing.HashRing
}

func NewCacheManager(hr *hashing.HashRing) CacheManager {
	return &cacheManager{hashRing: hr}
}

func (c *cacheManager) GetRedisClient(key string) rds.CacheClient {
	node := c.hashRing.GetNode(key)
	fmt.Printf("key: %s - node: %s\n", key, node.Title)
	rc := c.hashRing.Clients[node.Address]
	return rc
}
