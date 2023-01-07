package main

import (
	"fmt"
)

type person struct {
	name   string
	age    uint8
	height uint16
}

type student struct {
	person
	program    string
	university string
}

func main() {
	fmt.Println("hierarchy")
	joe := student{person{"Joe", 23, 180}, "computer Science", "mit"}
	fmt.Println(joe)
}
