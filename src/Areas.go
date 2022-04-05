package main

import (
	"fmt"
	"math"
)

func AreaCuadrado(base float64) float64 {
	return base * base
}

func AreaRectangulo(base float64, altura float64) float64 {
	return base * altura
}

func AreaCirculo(pi float64, radio float64) float64 {
	return pi * math.Pow(radio, 2)
}

func AreaTrapecio(b1 float64, b2 float64, h float64) float64 {
	return ((b1 + b2) * h) / 2
}

func main() {
	fmt.Println("El area del cuadrado: ", AreaCuadrado(10))
	fmt.Println("El area del rectangulo: ", AreaRectangulo(10, 5))
	fmt.Println("El area del circulo ", AreaCirculo(3.14, 5))
	fmt.Println("El area del trapecio: ", AreaTrapecio(6, 8, 4))

}
