package constant

import (
	"official-account-dev/cache"
	"strings"
)

type CacheKey int32

const (
	CacheAPIKey CacheKey = iota
	CacheAPISecret
	CacheWenXinToken
	CacheWenXinMsg
)

func GetCacheKey(key CacheKey) string {
	switch key {
	case CacheAPIKey:
		return strings.Join([]string{cache.Prefix, "CacheAPIKey"}, cache.Pattern)
	case CacheAPISecret:
		return strings.Join([]string{cache.Prefix, "CacheAPISecret"}, cache.Pattern)
	case CacheWenXinToken:
		return strings.Join([]string{cache.Prefix, "CacheWenXinToken"}, cache.Pattern)
	case CacheWenXinMsg:
		// %s 为FromUserName
		return strings.Join([]string{cache.Prefix, "CacheWenXinMsg", "%s"}, cache.Pattern)
	}
	return ""
}
