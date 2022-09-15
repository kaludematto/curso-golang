package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

// Só pode ter um parametro variatico e tem que ser o ultimo declarado
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo", 10, 20, 30)
}
