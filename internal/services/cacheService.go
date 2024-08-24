package services

import (
	"GorillaCacher/pkg/cache"
	"GorillaCacher/pkg/hashing"
	"time"
)

type CacheService interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, ttl time.Duration) error
}

type CacheServiceImpl struct {
	cacheManager cache.CacheManager
}

func NewCacheService(hr *hashing.HashRing) CacheService {
	return &CacheServiceImpl{cacheManager: cache.NewCacheManager(hr)}
}

func (cs *CacheServiceImpl) Get(key string) (interface{}, error) {
	rdc := cs.cacheManager.GetRedisClient(key)
	v := rdc.Get(key)
	return v, nil
}

func (cs *CacheServiceImpl) Set(key string, value interface{}, ttl time.Duration) error {
	rdc := cs.cacheManager.GetRedisClient(key)
	rdc.Set(key, value, ttl)
	return nil
}
