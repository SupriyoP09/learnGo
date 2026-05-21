package main

import (
	"fmt"
)

// channels are a way to communicate between goroutines.
// They are like pipes that connect two goroutines and allow them to send and receive values of a specific type.

// send
// func processNum(numChan chan int) {

// 	for num := range numChan {
// 		fmt.Println("processing number", num)
// 		time.Sleep(time.Second)
// 	}
// } // This function will be run as a goroutine and will process numbers received from the channel.

// func main() {

// rand.Seed(time.Now().UnixNano()) // Seed the random number generator to get different numbers each time we run the program.

// numChan := make(chan int)

// go processNum(numChan)

// for {
// 	numChan <- rand.Intn(100) // Send a random number between 0 and 99 to the channel. This will block until the processNum goroutine receives the number.
// }

// messagesChan := make(chan string) // channel of type string

// messagesChan <- "ping" // send a value to the channel // blocking operation, the main goroutine will wait until another goroutine receives the value

// msg := <-messagesChan // receive a value from the channel

// fmt.Println(msg)
// }

// receive
// func sum(results chan int, nums1 int, nums2 int) {
// 	numResults := nums1 + nums2
// 	results <- numResults // send the result to the channel
// }

// func main() {
// 	results := make(chan int)

// 	go sum(results, 1, 8)
// 	res := <-results // receive the result from the channel // blocking

// 	fmt.Println(res)
// }

// gorutine synchronization
// func task(done chan bool) {
// 	defer func() {
// 		done <-true
// 	} ()
// 	fmt.Println("processing...")
// }

// func main() {
// 	done := make(chan bool)
// 	go task(done)

// 	<-done // block
// }

// buffered channels
// func emailSender(emailChan chan string, done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	for email := range emailChan {
// 		fmt.Println("sending email to", email)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	emailChan := make(chan string, 5) // create a buffered channel with a capacity of 5
// 	done := make(chan bool)

// 	go emailSender(emailChan, done)

// 	for i:=0; i<5; i++ {
// 		emailChan <- fmt.Sprintf("%d@gmail.com", i) // send 5 emails to the channel
// 	}

// 	fmt.Println("done sending")

// fatal error: all goroutines are asleep - deadlock!

// daedlock fix
// close(emailChan)

// emailChan <- "1@email.com" // send the first email to the channel
// emailChan <- "2@gmail.com" // send the second email to the channel

// fmt.Println(<-emailChan) // receive the first email from the channel
// fmt.Println(<-emailChan) // receive the second email from the channel

// 	<-done // block until the emailSender goroutine is done
// }

// select statement
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received from chan2", chan2Val)
		}
	}

}
