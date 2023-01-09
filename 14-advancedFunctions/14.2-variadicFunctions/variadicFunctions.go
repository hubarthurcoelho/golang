package main

import "fmt"

// you can pass any number of arguments to a variadic function using ...
func sumAllNumbers(nums ...int) (total int) {
	for _, value := range nums {
		total += value
	}
	return
}

func main() {
	fmt.Println("variadic functions")
	total1 := sumAllNumbers(1, 2, 3, 4, 5)
	total2 := sumAllNumbers(2, 3, 4)
	fmt.Println(total1, total2)
}
