package main

import "fmt"

// defer delays the execution of a piece of code to the last possible call
func main() {
	fmt.Println("initializing main...")
	defer fmt.Println("main finished successfully")

	fmt.Println("doing stuff...")
}
