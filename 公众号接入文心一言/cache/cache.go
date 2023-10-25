package cache

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

const (
	Prefix  = "cache"
	Pattern = ":"
)

var (
	localCache *cache.Cache
	initOnce   sync.Once
)

func InitCache() {
	initOnce.Do(func() {
		localCache = cache.New(5*time.Minute, 10*time.Minute) // 设置默认过期时间为5分钟，清理过期项间隔为10分钟
	})
}

func Get(key string) (interface{}, bool) {
	return localCache.Get(key)
}

func Set(key string, value interface{}, d time.Duration) {
	localCache.Set(key, value, d)
}

func Delete(key string) {
	localCache.Delete(key)
}
