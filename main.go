package main

import (
	"MyStackoverflow/cache"
)

func main() {
	r := RegisterRouter()
	// pre-computed cache
	cache.Init()
	// listen and serve on 0.0.0.0:8080
	err := r.Run()
	if err != nil {
		return
	}
}
