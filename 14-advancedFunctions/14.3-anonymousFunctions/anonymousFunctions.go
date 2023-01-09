package main

import "fmt"

func main() {
	fmt.Println("anonymous function")
	func(text string) {
		fmt.Println(text)
	}("this is an anonymous function")
}
