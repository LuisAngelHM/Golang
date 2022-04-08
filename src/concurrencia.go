package main

//La concurrencia está lideando con multiples cosas al mismo tiempo
//Paralelismo: Está haciendo multiples cosas al mismo tiempo

//La función manin se ejecuta en un goroutine

import (
	"fmt"
	"sync"
)

func think(text string, wg *sync.WaitGroup) {
	//recibe un puntero de tipo waitgroup
	defer wg.Done()
	//defer funciona para ejecutar la instrucción hasta el utlimo.
	//wg.Done() termina con el hilo de ejcucion

	fmt.Println("I am thinking about ", text)

}

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {

	var wg sync.WaitGroup
	//se crea un WaitGroup para agrupar un conjunto de goroutines
	// no se utilizan mucho
	fmt.Println("Hello")
	wg.Add(1)
	//se agrega una gorountine al WaitGroup
	go say("World", &wg)
	wg.Add(1)

	go think("You", &wg)

	wg.Wait()
	//Espera a que terminen los diferentes goroutine
}
