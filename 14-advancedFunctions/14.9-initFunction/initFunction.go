package main

import "fmt"

func main() {
	fmt.Println("main")
}

// the init function is called before the main function.
// you can have one init function per file
func init() {
	fmt.Println("initializing...")
}
