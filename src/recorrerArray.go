package main

import (
	"fmt"
	"strings"
)

func isPalindromo(dta string) bool {
	dta = strings.ToLower(dta)
	var textReverse string
	for i := len(dta) - 1; i >= 0; i-- {
		textReverse += string(dta[i])
	}

	if textReverse == dta {
		return true
	}
	return false
}
func main() {
	slice := []string{"hola", "que", "hace"}
	for _, valor := range slice {
		fmt.Println(valor)
	}

	fmt.Println(isPalindromo("amor a roma"))
}
