package db

import (
	"aegis_test/libs"
	"aegis_test/libs/custom_error"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type CacheDB interface {
	SetCache(key string, obj interface{}) error
	GetCache(key string, target any) (interface{}, error)
	RemoveCacheForContext(contextName string)
	CacheKeyBuilder(contextName string, params ...string) string
}

type cacheDB struct {
	redisClient  *redis.Client
	cacheClient  *cache.Cache
	cacheContext context.Context
}

func InitCache() *cacheDB {
	cdb := cacheDB{}

	log.Println("Trying to Connect to Cache Server...")
	cdb.redisClient = redis.NewClient(&redis.Options{
		Addr:     libs.RedisHost + ":" + libs.RedisPort,
		Password: libs.RedisPassword,
		DB:       libs.RedisDB,
	})

	cdb.cacheClient = cache.New(&cache.Options{
		Redis:      cdb.redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	cdb.cacheContext = context.TODO()
	log.Println("Connected to Cache Server Successfully!")

	return &cdb
}

func (cdb *cacheDB) SetCache(key string, obj interface{}) error {
	if cdb.cacheClient != nil {
		if json, err := json.Marshal(obj); err == nil {
			cdb.cacheClient.Set(&cache.Item{
				Ctx:   cdb.cacheContext,
				Key:   key,
				Value: &json,
				TTL:   time.Hour,
			})
			return nil
		}
	}

	return custom_error.UnavailableService()
}

func (cdb *cacheDB) GetCache(key string, target any) (interface{}, error) {
	if cdb.cacheClient != nil {
		var wanted interface{}
		if err := cdb.cacheClient.Get(cdb.cacheContext, key, &wanted); err == nil {
			err := json.Unmarshal(wanted.([]byte), &target)
			if err == nil {
				return target, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return nil, custom_error.UnavailableService()
}

func (cdb *cacheDB) RemoveCacheForContext(contextName string) {
	if cdb.redisClient != nil {
		pattern := fmt.Sprintf("%s::*", cdb.CacheKeyBuilder(contextName))
		i := cdb.redisClient.Scan(cdb.cacheContext, 0, pattern, 0).Iterator()
		if i.Err() == nil {
			for i.Next(cdb.cacheContext) {
				cdb.redisClient.Del(cdb.cacheContext, i.Val())
			}
		} else {
			log.Println("Failed clean cache keys", pattern)
		}
	}
}

func (cdb *cacheDB) CacheKeyBuilder(contextName string, params ...string) string {
	keyName := fmt.Sprintf("%s::%s", libs.CacheKeyPrefix, contextName)

	for _, param := range params {
		keyName = fmt.Sprintf("%s::%s", keyName, param)
	}

	return keyName
}
