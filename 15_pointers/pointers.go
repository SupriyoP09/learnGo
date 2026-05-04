package main

import "fmt"

// when we pass a variable to a function, 
// it creates a copy of that variable and passes it to the function. 
// So, any changes made to the variable inside the function do not affect the original variable outside the function.

// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In changeNum", num)
// }

// func main() {
// 	num := 1

// 	changeNum(num)
// 	fmt.Println("after changeNum main", num)
// }

// by reference
// when we pass a variable to a function, it creates a copy of that variable and passes it to the function.
func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num := 1
	changeNum(&num)
	// fmt.Println("Memory address", &num) // operator to get the memory address of a variable.
	fmt.Println("after changeNum main", num)
}