package main

import "fmt"

func sum(nums ...int) int {
	total := 0

	for _, nums := range nums {
		total = total + nums
	}

	return total
}

func main() {

	nums := []int{1, 2, 3, 4, 8}
	result := sum(nums...)
	
	// result := sum(1, 2, 3, 4, 8)
	fmt.Println(result)
}