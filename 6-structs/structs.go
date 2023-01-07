// structs are like interfaces in typescript
package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint16
}

func main() {
	var arthur user = user{name: "Arthur", age: 23}
	fmt.Println(arthur)
	devis := user{"Devis", 31, address{"cond", 17}}
	fmt.Println(devis)
	lucas := user{name: "Lucas", age: 23}
	lucasAddress := address{"Cond", 17}
	lucas.address = lucasAddress
	fmt.Println(lucas)
}
