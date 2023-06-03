package main

import (
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func main() {
	println("started")
	time.Sleep(1 * time.Second)
	go Raftperiodic()
	time.Sleep(5 * time.Second)
	mu.Lock()
	done = true
	mu.Unlock()
	println("exiting")
	time.Sleep(3 * time.Second)
}
func Raftperiodic() {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			return
		}
		mu.Unlock()
	}

}
