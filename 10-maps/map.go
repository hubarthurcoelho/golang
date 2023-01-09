package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// ! in a map, keys should have the same type, as well as values
	user := map[string]string{
		"name":    "Arthur",
		"surname": "Coelho",
	}

	users := map[string]map[string]string{
		"user1": {
			"name":    "Devis",
			"surname": "Coelho",
		},
		"user2": {
			"name":    "Lucas",
			"surname": "Coelho",
		},
	}
	fmt.Println(user["name"])
	fmt.Println(users)
	// deleting a map key
	delete(users, "user2")
	fmt.Println(users)
	// creating a key
	users["user2"] = map[string]string{
		"name":    "Matheus",
		"surname": "Coelho",
	}
	fmt.Println(users)
}
