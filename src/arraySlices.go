package main

import "fmt"

func main() {
	//arrays
	//son inmutbles
	var array [4]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array, len(array), cap(array))

	//slices
	//son mutables
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	//metodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	//Append

	slice = append(slice, 7)
	fmt.Println(slice)
	//Append Lits
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

	newSlice = []int{}
	fmt.Println(newSlice)
	newSlice = append(newSlice, 5)
	fmt.Println(newSlice)	
}
