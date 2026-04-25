package main

import (
	"fmt"
	//"time"
)

func main() {
	// simple switch
	//i := 2

	// u dont need to break in go, it automatically breaks after each case
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")

	// default:
	// 	fmt.Println("other")
	// }

	// switch with multiple cases

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's workday")
	// }

	// type switch
	whoAmI := func(i interface{}) {
		switch t:= i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("its a boolean")
		default:
			fmt.Printf("other", t)
		}
	}

	//whoAmI("golang")
	whoAmI(18)
	// whoAmI(true)
	// whoAmI(3.14)
}