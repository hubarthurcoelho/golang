package main

import "fmt"

// when using named return values, you can declare the variable in the return type
// and when returning, you dont have to return the variables. Just return.
func calculate(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	fmt.Println("named Return values")
	sum, sub := calculate(1, 2)
	fmt.Println(sum, sub)
}
