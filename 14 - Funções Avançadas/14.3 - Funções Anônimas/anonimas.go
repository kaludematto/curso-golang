package main

import "fmt"

func main() {

	func() { // Declaração da função anônima
		fmt.Println("Olá Mundo")
	}() //necessário adicionar no final da função para ela ser chamada anônimamente

	// Função Anônima com parametros
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Olá Mundo de novo...") // parametros passados no final da função
	fmt.Println(retorno)
}
