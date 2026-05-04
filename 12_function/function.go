package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string, bool) {
	return "golang", "kotlin", "python", true
}

// func processIt(fn func(a int) int) {
// 	fn(2)
// }

// func processIt() func(a int) int {
// 	return func (a int) int {
// 		return 4
// 	}
// }

func main() {
	// fn := func(a int) int {
	// 	return 2
	// }
	// fn := processIt()
	// fn(6)


	lang1, lang2, lang3, _ := getLanguages()
	fmt.Println(lang1, lang2, lang3)
	result := add(3, 5)
	fmt.Println(result)
}