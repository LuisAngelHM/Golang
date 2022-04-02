package main

import "fmt"

func main() {
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma", result)

	result = y - x
	fmt.Println("Resta: ", result)

	result = x * y
	fmt.Println("Multiplicacion: ", result)

	result = y / x
	fmt.Println("Division: ", result)

	result = y % x
	fmt.Println("Modulo", result)

	x++
	fmt.Println("incremental: ", x)

	x--
	fmt.Println("Decremental: ", x)
}
