package main

import (
	"fmt"
)

func main() {
	//Defer: ejecuta al ultimo la linea que se especificó
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue continu el ciclo for
	//break rompe el ciclo

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es dos")
			continue
		} else if i == 5 {
			break
		}
	}
}
