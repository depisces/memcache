package cache_server

import (
	"memcache/cache"
	"time"
)

type cacheServer struct {
	memCache cache.Cache
}

func NewMemCache() *cacheServer {
	return &cacheServer{
		memCache: cache.NewMemCache(),
	}
}

func (cs *cacheServer) SetMaxMemory(size string) bool {
	return cs.memCache.SetMaxMemory(size)
}
func (cs *cacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	expireTs := time.Second * 0
	if len(expire) > 0 {
		expireTs = expire[0]
	}

	return cs.memCache.Set(key, val, expireTs)
}

// 根据key值获取value
func (cs *cacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}

// 删除key值
func (cs *cacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

// 判断key是否存在
func (cs *cacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

// 清空所有key
func (cs *cacheServer) Flush() bool {
	return cs.memCache.Flush()
}

// 获取缓存中key的数量
func (cs *cacheServer) Keys() int64 {
	return cs.memCache.Keys()
}
