package main

import "fmt"

//funcion que retorna dos valores
func retornaDosValores(a int) (c, d int) {

	return a * 2, a * 4
}

func main() {
	helloMessage := "Helo"
	worldMessage := "World"

	//Println: Agrega un salto de linea
	fmt.Println(helloMessage, worldMessage)

	//Printf: Agrega una funcion extra al string de  entrada
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n ", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n ", nombre, cursos)

	//Sprintf: No imprime en consola, solo guarda el string

	message := fmt.Sprintf("%s tiene más de %d cursos\n ", nombre, cursos)
	fmt.Println(message)

	//Conocer el tipo de dato de una variable

	fmt.Printf("helloMessage: %T \n", helloMessage)
	fmt.Printf("Cursos: %T\n", cursos)

	// si solo se quiere ocupar uno de los dos valores se debe utiliza '_'
	value1, _ := retornaDosValores(5)
	fmt.Println(value1)

}
