package main

import (
	"fmt"
	"sync"
)

type post struct {
	view int
	mu sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func ()  {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.view++
}

func main() {
	var wg sync.WaitGroup
	myPost := post{view: 0}

	for i := 0; i<100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)

	}

	wg.Wait()
	fmt.Println(myPost.view)
}

// 1. Mutex is a mutual exclusion lock that can be used to protect shared resources from concurrent access.
// 2. A mutex can be used to ensure that only one goroutine can access a shared resource at a time.
// 3. The sync.Mutex type provides two methods: Lock() and Unlock().
// 4. The Lock() method is used to acquire the lock, and the Unlock() method is used to release the lock.
// 5. It is important to always release the lock after acquiring it, otherwise other goroutines will be blocked indefinitely.