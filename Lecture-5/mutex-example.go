package main

import (
	"sync"
	"time"
)

func main() {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter += 1

		}()
	}
	time.Sleep(1 * time.Second)
	println(counter)

}
