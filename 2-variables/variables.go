package main

import "fmt"

func main() {
	var variable1 string = "variable 1"
	fmt.Println(variable1)
	variable2 := "variable 2"
	fmt.Println(variable2)
	var (
		variable3 string = "variable 3"
		variable4 string = "variable 4"
	)
	fmt.Println(variable3, variable4)

	variable5, variable6 := "variable 5", "variable 6"
	fmt.Println(variable5, variable6)
	variable5, variable6 = variable6, variable5
	fmt.Println(variable5, variable6)
}
