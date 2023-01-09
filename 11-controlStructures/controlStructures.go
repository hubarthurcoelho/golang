package main

import "fmt"

func main() {
	fmt.Println("control structures")

	numero := 10
	if numero > 15 {
		fmt.Println("greater than 15")
	} else {
		fmt.Println("less than or equal to 15")
	}

	// if init (creating a variable inside an if statement. local scope only)
	if number := numero; number > 0 {
		fmt.Println("greater than 0")
	}
}
