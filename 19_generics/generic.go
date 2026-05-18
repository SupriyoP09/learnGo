package main

import (
	"fmt"
)

func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }


func main() {
	// nums := []int{18,7,45}
	// names := []string{"Virat", "Dhoni", "Rohit"}
	vals := []bool{true, false, true}
	printSlice(vals)
	// printSlice(names)
}

// LIFO
// type stack[T any] struct {
// 	elements []T
// }

// func main() {
// 	myStack := stack[string]{
// 		elements: []string{"Virat"},
// 	}

// 	fmt.Println(myStack)
// }