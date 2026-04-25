package main

import "fmt"

func main() {
	// age := 14

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is a not an adult")
	// }

	// age := 16

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else if age >= 12 {
	// 	fmt.Println("person is a teenager")
	// } else {
	// 	fmt.Println("person is a child")
	// }

	// var role = "admin"
	// var hasPermission = false
	// var hasPermission = true

// 	if role == "admin" && hasPermission {
// 		fmt.Println("Access granted")
	// }

    // we can also declare and initialize a variable in the if statement itself
	if age :=20; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is a teenager", age)
	}

	// go does not have a ternary operator like in other languages, but we can achieve similar functionality using if-else statements
	
}