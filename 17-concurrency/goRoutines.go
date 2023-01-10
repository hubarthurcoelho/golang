package main

import (
	"fmt"
	"time"
)

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}

func main() {
	// write("Hello, world!")     this would run forever and never call the second line
	// write("please read me...")  but we can use go routines to make it run concurrently
	go write("Hello, world!")  // using go, this will run "in the background"
	write("please read me...") // if we put go here, the function ends and nothing happens
}
