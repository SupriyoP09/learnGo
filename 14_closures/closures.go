package main

import "fmt"

// A closure is a function that captures the variables from its surrounding scope. 
// In Go, closures are created when you define a function inside another function 
// and the inner function references variables from the outer function.
func counter() func() int {
	var count int = 0
// The inner function is a closure that captures the `count` variable from the outer function's scope.	
	return func() int {
		count += 1
		return count
	}
}
// In this example, the `counter` function returns an anonymous function that increments and returns the `count` variable.
func main() {
	// When we call `counter()`, it returns the inner function, which is a closure that has access to the `count` variable. 
	// Each time we call the returned function, it increments and returns the updated value of `count`.
	increment := counter()

	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
}