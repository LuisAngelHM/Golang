//Structs sirve para realizar clases
package main

import "fmt"

type car struct {
	//para definir si es publico, la variable debe de empezar con la letra Mayuscula
	//para definir si es privado, la varibale debe de empezar con la letra minucula
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	var otherCar car
	otherCar.brand = "Ferrari"
	otherCar.year = 12555
	fmt.Println(otherCar)
}
