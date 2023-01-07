package main

import "fmt"

// points to a memory address in your computer
// to use a pointer, use * before the type declaration
// use & before the value atribution
func main() {
	var num1 int = 10
	var num1Pointer *int = &num1
	fmt.Println(*num1Pointer)
}

// to get the value instead of the address, use * before the variable call
