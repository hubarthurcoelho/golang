// operators
// +, -, /, %, x++, x--, x +,-,/...= y
// >, >=, <, <=, ==, !=
// theres no ternary, use if else instead
// &&, ||
package main

import "fmt"

func main() {
	sum := 1 + 1
	sub := 1 - 1
	mult := 1 * 1
	div := 1 / 1
	rest := 2 % 3
	if sum > sub {
		fmt.Println("sum > sub")
	} else {
		fmt.Println("sub > sum")
	}
	if mult == div && rest > div {
		fmt.Println("mult == div && rest > div")
	}
}
