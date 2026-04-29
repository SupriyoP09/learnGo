package main

import (
	"fmt"
)

// interating over a range of numbers
func main() {
	// nums := []int {1,2,3}

	// one way to iterate over a range of numbers is using a for loop
	// for i:=0; i < len(nums); i++ {
	// 	fmt.Println(nums[i])
	// }

	// other way to iterate over a range of numbers is using the range keyword
	// for _, num := range nums {
	// 	fmt.Println(num)
	// }


	// we can also use the range keyword to sum up a range of numbers
	// sum := 0
	// for _, num := range nums {
	// 	sum += num
	// }
	// fmt.Println(sum)
	// fmt.Println("sum:", sum)

	// value and index
	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	// m := map[string]string{"fname": "John", "lname": "Doe"}

	// for k, v :=range m {
	// 	fmt.Println(k, v)
	// }

	// for k:=range m {
	// 	fmt.Println(k)
	// }

	// unicode codes point rune
	// starting byte of rune
	// 300 ->  -> 1 byte, 2 bytes
	for i, c:=range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}

}