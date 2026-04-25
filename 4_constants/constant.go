package main

import "fmt"

// const age = 30

func main() {
	// const name = "samosa"
	// const age = 30

	// fmt.Println(age)

	const (
		port = 8080
		host = "localhost"
	)

	fmt.Println(port,host)
}