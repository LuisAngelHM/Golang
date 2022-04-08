//Interfaces: Una interfaz es un método que comparte diferentes métodos de diferentes estructuras.

package main

import "fmt"

type figuras2D interface {
	area() float64
}

type rectangulo struct {
	base   float64
	altura float64
}

type cuadrado struct {
	base float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculate(f figuras2D) {

	fmt.Println("Área: ", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 5}

	calculate(myCuadrado)
	calculate(myRectangulo)

	//lista de interfaces: Pueden almacenar diferentes tipo de datos
	myInterface := []interface{}{"Hola", 12, 4.5}
	fmt.Println(myInterface)
}
