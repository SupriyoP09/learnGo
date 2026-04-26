package main

import (
	"fmt"
	"maps"
)

// maps -> hash, objects, dict

func main() {
	// creating a mapp

	// m := make(map[string]string)

	// setting a element
	// m["name"] = "go"

	// get an element
	// fmt.Println(m["name"], m["area"])

	// if the key is not present, it will return the zero value
	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(m["phone"])

	// fmt.Println(len(m))

	// delete an element
	// delete(m, "price")

	// clear
	// clear(m)
	// fmt.Println(m)


	// m := map[string]int{
	// 	"price": 50,
	// 	"phone": 30,
	// }
	
	// v, ok := m["phone"]
	// fmt.Println(v)

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	m1 := map[string]int{
		"price": 50,
		"phone": 30,
	}	

	m2 := map[string]int{
		"price": 50,
		"phone": 30,
	}

	fmt.Println(maps.Equal(m1,m2))

}