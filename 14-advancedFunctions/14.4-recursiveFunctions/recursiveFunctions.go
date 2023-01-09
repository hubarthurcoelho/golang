package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return num
	} else {
		num *= factorial(num - 1)
	}
	return num
}

func main() {
	fmt.Println("recursive functions")
	factorial := factorial(5)
	fmt.Println(factorial)
}
