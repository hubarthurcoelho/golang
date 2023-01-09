package main

import "fmt"

func genericFunc(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("generics")
	genericFunc(1)
	genericFunc("string")
	genericFunc(true)
}
