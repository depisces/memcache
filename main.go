package main

import (
	"fmt"
	cache_server "memcache/cache-server"
	"time"
)

func main() {
	//内存缓存系统
	cache := cache_server.NewMemCache()
	cache.SetMaxMemory("200MB")

	cache.Set("int", 1, time.Second)
	cache.Set("bool", false, time.Second)
	cache.Set("data", map[string]interface{}{"a": 1}, time.Second)
	cache.Set("nihao", "zheshishenme")
	cache.Del("int")
	cache.Exists("key")
	cache.Keys()
	fmt.Println(cache.Get("niha"))
	fmt.Printf("cache.Keys(): %v\n", cache.Keys())
}
