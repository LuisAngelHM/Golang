package main

import "fmt"

func main() {

	//for condicional

	for i := 1; i < 11; i++ {
		//Tablas de multiplicar
		fmt.Printf("%d x %d = %d \n", 5, i, (5 * i))
	}

	//for while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	// For forever: Nunca termina

	counForever := 0
	for {
		fmt.Println(counForever)
		counForever++
	}

}
