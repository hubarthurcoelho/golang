package main

import "fmt"

func invertNumberSign(num int) int {
	return num * -1
}

func invertNumberSignPointer(num *int) {
	*num = *num * -1
}

func main() {
	num1 := 20
	// doing it this way, the num1 is just a copy of the variable above
	newNum := invertNumberSign(num1)
	fmt.Println(num1, newNum)

	num2 := 10
	fmt.Println(num2)
	invertNumberSignPointer(&num2)
	fmt.Println(num2)
}
