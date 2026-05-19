package main

import (
	"fmt"
	// "time"
	"sync"
)

// func task(id int) {
// 	fmt.Println("doing task", id)
// }

// func main() {
// 	for i := 0; i <= 10; i++ {
// 		// go task(i)

// 		go func(i int) {
// 			fmt.Println(i)
// 		}(i)
// 	}

// 	time.Sleep(time.Second * 2)
// }

// WaitGroup

// WaitGroup is a synchronization primitive that can be used to wait for a collection of goroutines to finish.
func task(id int, w *sync.WaitGroup) {
	defer w.Done() // defer function will be executed at the end of the function, even if there is an error
	fmt.Println("doing task", id)
}

func main() {
	var wg sync.WaitGroup // create a WaitGroup to wait for all goroutines to finish

	for i := 0; i <= 10; i++ {
		wg.Add(1) // minus one from the WaitGroup counter for each goroutine, and when the counter reaches zero, the WaitGroup will unblock
		go task(i, &wg)
	}

	wg.Wait() // wait for all goroutines to finish
}
