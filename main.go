package main

import (
	cacheGo "Klivitskyi/go-cache/cache-go"
	"fmt"
	"time"
)

func main() {
	cache := cacheGo.New()

	fmt.Print(cache.Set("userId", 42, time.Second))

	time.Sleep(time.Second * 2)

	userId, ok := cache.Get("userId")

	if !ok {
		fmt.Printf("The user with ID %v didn't find\n", userId)
		return
	}

	fmt.Println(userId)

	cache.Delete("userId")
	userId, ok = cache.Get("userId")

	fmt.Println(userId, ok)
}
