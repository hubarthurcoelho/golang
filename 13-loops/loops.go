package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("i", i)
		time.Sleep(time.Millisecond) // pauses the code for some time
	}

	// declaring inside for scope
	for j := 0; j < 10; j++ {
		fmt.Println("j", j)
		time.Sleep(time.Millisecond)
	}

	// using for and range to loop over an array
	array := []string{"Arthur", "Devis", "Lucas"}
	for index, value := range array {
		fmt.Println(index, value)
	}

	// you can use an _ if you dont want the index or the value
	for _, value := range "Word" {
		// to get the actual letter, you need to use string()
		fmt.Println(string(value))
	}

	namesMap := map[string]map[string]string{
		"Arthur": {
			"surname": "Coelho",
		},
		"Devis": {
			"surname": "Coelho",
		},
	}

	for index, value := range namesMap {
		fmt.Println(index, value["surname"])
	}

}
