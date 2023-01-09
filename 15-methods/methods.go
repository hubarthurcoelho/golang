package main

import "fmt"

type user struct {
	name    string
	surname string
	age     int
}

func (u user) save() {
	fmt.Printf("saving user %s\n", u.name)
}

// you have to pass user as a pointer to actually change the users age
func (u *user) birthday() {
	u.age++
}

func main() {
	arthur := user{"Arhur", "Coelho", 23}
	arthur.save()
	arthur.birthday()
	fmt.Println(arthur.age)
}
