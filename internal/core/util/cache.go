package util

import (
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/sirupsen/logrus"
)

var gloabalCache *ristretto.Cache

const (
	CacheNumCounters = 1000
	CacheMaxCost     = 1 << 30
	CacheBufferItems = 64
	CacheExpiration  = 2 * time.Minute
)

func GetCache() *ristretto.Cache {
	if gloabalCache == nil {
		cache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: CacheNumCounters,
			MaxCost:     CacheMaxCost,
			BufferItems: CacheBufferItems,
		})

		if err != nil {
			panic("Cannot setup cache.")
		}
		gloabalCache = cache
	}

	return gloabalCache
}

func AddtoCache(key string, value interface{}) {
	if gloabalCache == nil {
		GetCache()
	}
	gloabalCache.Set(key, value, 0)
	ExpireCache(key)
}

func ExpireCache(key string) {
	go func() {
		time.Sleep(CacheExpiration)
		logrus.Infof("Expiring entry: %s", key)
		gloabalCache.Del(key)
	}()
}
