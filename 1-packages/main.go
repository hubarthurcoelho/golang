package main

import (
	"fmt"
	"modules/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main")
	helper.Helper()
	erro := checkmail.ValidateFormat("arthur.coelho@gmail.com")
	fmt.Println(erro)
}
