package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var numero int16 = 100
	fmt.Println(numero)

	var numero2 uint32 = 120
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12345
	fmt.Println(numero3)

	//BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// NUMEROS REAIS

	var numero5 float32 = 3.25
	fmt.Println(numero5)

	var numero6 float64 = 3.155555555
	fmt.Println(numero6)

	numero7 := 4.555555
	fmt.Println(numero7)

	// FIM NUMEROS REAIS

	// COMEÇO STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := '8'       // Sempre que usar aspas simples o go pega o numero da tabela ASCII
	fmt.Println(char) // Imprime o numero da tabela ASCII que representa o '8'

	// FIM STRINGS

	var texto string
	fmt.Println(texto)

	//BOOLEAN
	verdadeiro := true
	fmt.Println(verdadeiro)
	falso := false
	fmt.Println(falso)

	// ERRO
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)
}
