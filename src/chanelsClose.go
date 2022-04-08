package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	fmt.Println(len(c), cap(c))

	//Range:recorre todos los elementos de array, slice, map o canales
	//Close: Le dice a goroutine cuando cerrar el canal, lo ideal es cerrar los canales cuando se dejan de usar
	close(c)

	for message := range c {
		fmt.Println(message)
	}

	//Select: elije un canal obtener el dato
	email1 := make(chan string)
	email2 := make(chan string)
	go message("mensaje1", email1)
	go message("mensaje2", email2)

	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
