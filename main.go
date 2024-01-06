package main

import (
	cacheGo "Klivitskyi/go-cache/cache-go"
	"fmt"
)

func main() {
	cache := cacheGo.New()

	fmt.Println(cache.Set("userId", 42))
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
