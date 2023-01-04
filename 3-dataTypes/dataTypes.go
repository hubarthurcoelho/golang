package main

import (
	"errors"
	"fmt"
)

// int8(byte), int16, int32(rune - ASCII), int64, int
// float32, float64
// number of bits
func main() {
	var number int = 10000
	var number2 int = 20000
	var number3 uint = 1000 // unsigned int (only positive)
	fmt.Println(number, number2, number3)
	char := 'b'
	fmt.Println(char) // prints the equivalent number in the ASCII representation
	var boolean bool
	fmt.Println(boolean)
	var error error = errors.New("error")
	fmt.Println(error)
}
