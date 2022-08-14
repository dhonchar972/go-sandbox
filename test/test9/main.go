package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex // sync.RWMutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock() // c.mu.WLock()
	c.c[key]++
	c.mu.Unlock() // c.mu.WUnlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()         // c.mu.RLock()
	defer c.mu.Unlock() // c.mu.RUnlock()
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)} // create dictionary string: int
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value(key))

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	// some action
	//  atomic.Add(&count, 1)
	// }()

	var once sync.Once
	done := make(chan bool)
	for i := 0; i < 1000; i++ {
		go func() {
			once.Do(func() { // work only once
				fmt.Println("some action")
			})
			done <- true
		}()
	}
}
