package main

import "fmt"

type ApiOne struct {
	url   string
	rates []int
}

func (o ApiOne) getRates() []int {
	fmt.Println(o.url)
	return o.rates
}

type ApiTwo struct {
	url    string
	quotes []int
}

func (t ApiTwo) getRates() []int {
	fmt.Println(t.url)
	return t.quotes
}

type Api interface {
	getRates() []int
}

func getRates(a Api) []int {
	response := a.getRates()
	return response
}

type Person struct {
	name    string
	surname string
}

type IUser struct {
	Person
	age uint8
}

func (u IUser) getName() string {
	return u.name
}

func double(num *int) {
	*num = *num * 2
}

func namedReturnValue(num int) (result int) {
	result = num * 2
	return
}

// can accept other values, but only before the spread
func variadic(num ...int) (result int) {
	for _, value := range num {
		result += value
	}
	return
}

func anonymous(num int) {
	func(param *int) {
		*param *= 2
	}(&num)
}

func recursiveFunction(num int) int {
	if num == 1 {
		return num
	} else {
		num *= recursiveFunction(num - 1)
		return num
	}
}

func callingDefer(text string) {
	fmt.Println("Initializing calling defer...")
	defer fmt.Println("calling Defer finished")

	fmt.Println(text)
}

func main() {
	arthur := IUser{Person{"Arthur", "Coelho"}, 23}
	fmt.Println(arthur)
	num1 := 10
	fmt.Println(num1)
	double(&num1)
	fmt.Println(num1)
	array := [3]int{1, 2, 3}
	fmt.Println(array, len(array), cap(array))
	slice := []int{}
	fmt.Println(slice, len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))
	mapa := map[string]IUser{
		"arthurcoelho@gmail.com": {
			Person{"Arthur", "Coelho"},
			23,
		},
	}
	fmt.Println(mapa)

	// LOOPS

	slicers := []string{"arthur", "lucas", "matheus", "devis"}
	for index, value := range slicers {
		fmt.Println(index, value)
	}
}
