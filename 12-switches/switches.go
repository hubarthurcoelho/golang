package main

import "fmt"

func main() {
	fmt.Println("switches")
	name := "Arthur"
	switch name {
	case "Arthur":
		fmt.Println("tuzao")
	case "Devis":
		fmt.Println("Devao")
	default:
		fmt.Println("Teuzao")
	}
	surname := "Coelho"
	switch surname {
	case "Coelho":
		fmt.Println("Coelho")
		fallthrough // this will read the next case block of code and execute it. Not used very often
	default:
		fmt.Println("Primos")
	}
}
