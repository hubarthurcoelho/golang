package main

import "fmt"

func closures() func() {
	text := "inside closure"

	function := func() {
		fmt.Println(text)
	}
	return function
}

func main() {
	fmt.Println("closures")
	newFunc := closures()
	newFunc() // prints inside closure
}
