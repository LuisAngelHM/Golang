package main

import "fmt"

//channels:Es un conducto en el cual solo puedes manejar un tipo de datos

func say(text string, c chan<- string) {
	c <- text
}

func main() {

	c := make(chan string, 1)
	fmt.Println("Hello")
	go say("bye", c)
	fmt.Println(<-c)
}
