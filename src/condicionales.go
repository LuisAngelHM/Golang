package main

import "fmt"

func main() {

	fmt.Println("Ingresa un numero: ")
	var dato int
	fmt.Scanln(&dato)

	if dato%2 == 0 {
		fmt.Println("El número", dato, "es par")
	} else {
		fmt.Println("El número", dato, "no es par")
	}

	fmt.Println("Ingresa un número de mes")
	fmt.Scanln(&dato)
	switch dato {
	case 1:
		fmt.Println("Enero")
		break
	case 2:
		fmt.Println("Febrero")
		break
	case 3:
		fmt.Println("Marzo")
		break
	case 4:
		fmt.Println("Abril")
		break
	case 5:
		fmt.Println("Mayo")
		break
	case 6:
		fmt.Println("Junio")
		break
	case 7:
		fmt.Println("Julio")
		break
	case 8:
		fmt.Println("Agosto")
		break
	case 9:
		fmt.Println("Septiembre")
		break
	case 10:
		fmt.Println("Octubre")
		break
	case 11:
		fmt.Println("Noviembre")
		break
	case 12:
		fmt.Println("Diciembre")
		break
	default:
		fmt.Println("No es un número valido")

	}

}
