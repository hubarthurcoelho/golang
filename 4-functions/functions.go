package main

import "fmt"

// functions need to have its params typed. The return should also be typed
// right after the params. Functions in go can return more than one value

// typing params
// func personInfo(name string, surname string, age int, height int) string {
// 	return "xd"
// }

// diffent way of typing params
// func personInfo2(name, surname string, age, height int) string {
// 	return "xd"
// }

// returning more than one value
func operations(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func main() {
	sum, sub := operations(1, 2)
	sum1, _ := operations(3, 1)
	// %d -> integer, %s -> string, %f -> floats https://pkg.go.dev/fmt
	fmt.Println(fmt.Sprintf("sum %d, sub %d, sum1 %d", sum, sub, sum1))
}
