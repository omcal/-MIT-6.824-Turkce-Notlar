package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	alice := 10000
	bob := 10000
	var mu sync.Mutex
	total := bob + alice
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			mu.Unlock()
			mu.Lock()
			bob += 1
			mu.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			mu.Unlock()
			mu.Lock()
			alice = +1
			mu.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) <= 1*time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("Hata alica =%v,bob=%v total=%v\n", alice, bob, total)

		}
		mu.Unlock()
	}

}
