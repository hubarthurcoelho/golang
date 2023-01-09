package main

import "fmt"

func recoverExecution() {
	// this will only run if recover() is != nil, so theres no problem calling it every time
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic")
	}
}

func averageStudentGrade(grade1, grade2 uint) bool {
	// this will be called when the function is executed
	defer recoverExecution()
	average := (grade1 + grade2) / 2
	if average > 6 {
		return true
	} else {
		// panic will stop the execution of your program
		// but you can use recover to keep it running
		panic("RIP")
		// keep in mind, panic is not the best way to treat errors.
	}
}
func main() {
	averageStudentGrade(6, 6)
	fmt.Println("hello world")

}
