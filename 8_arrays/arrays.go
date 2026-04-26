package main

import "fmt"

// numbered sequence of specfic length

func main() {

	// zero values
	// int -> 0
	// bool -> false
	// string -> ""
	// var nums [4]int

	// nums[0] = 1

	// fmt.Println(nums[0])
	// fmt.Println(nums)
	// array length
	// fmt.Println(len(nums))

	// var vals [4]bool
	// vals[2] = true
	// fmt.Println(vals) // [false false false false] // [false false true false]

	// var name [3]string
	// name[0] = "Virat"
	// fmt.Println(name) // [Virat  ]

	// to declare it in single line
	// nums := [4]int{1, 2, 3, 4}
	// fmt.Println(nums) // [1 2 3 4]

	// 2d array
	nums := [2][2]int{{4, 5},{1, 2}}
	fmt.Println(nums) // [[4 5] [1 2]]

	// - fixed size, that is predictable memory usage
	// - Memory optimazation, that is contiguous memory allocation
	// - Contant time access to elements
}