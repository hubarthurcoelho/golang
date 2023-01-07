package main

// go mod init "moduleName"

import (
	"fmt"
	"modules/functions"
	"modules/helper"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main")
	helper.Helper()
	erro := checkmail.ValidateFormat("arthur.coelho@gmail.com")
	fmt.Println(erro)
	sum := functions.Operations(1, 2)
	fmt.Println(sum)
}
