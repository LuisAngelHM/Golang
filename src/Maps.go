package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20
	fmt.Println(m)

	//recorrer map, por lo regular no se recorren en orden en el que se ingresaron los datos debido a la concurrencia
	for i, valor := range m {
		fmt.Println(i, valor)
	}

	valo1, ok := m["josep"]
	if ok == true {
		fmt.Println(valo1)
	}

}
