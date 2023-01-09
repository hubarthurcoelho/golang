package main

import "fmt"

func main() {
	// An Array has a fixed amount of elements, and they all have the same type
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	array[0] = 3 // to access a specific element
	// A Slice is a part of an array, and when its not provided, go uses an internal array
	fmt.Println(array)
	var slice []int = []int{1, 2, 3, 4, 5} // number of elements not fixed
	fmt.Println(slice)
	slice = append(slice, 6) // makes a copy of slice, appending an extra element
	fmt.Println(slice)
	// those examples were using slices from the internal array
	// you can also slice existing arrays
	var arraySlice []int = array[1:3] // inclusive:exclusive
	fmt.Println(arraySlice)
	// this slice will point to the memory address of the array elements. If the elements change, the sliced elements will also change
	array[1] = 4
	fmt.Println(arraySlice)
	// internal array
	slice1 := make([]int, 3, 5) // type : slice size : array length
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // len() gets the length of the slice
	fmt.Println(cap(slice1)) // cap() gets the maximum size of the array
	// if you append enough elements to surpass the array size, go will double the array max size
}
