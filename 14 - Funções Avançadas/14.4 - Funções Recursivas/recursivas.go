// Funções Recursivas são funções que chamam elas mesmas
// Para que ela funcione ela depende de outra função para chamar ela mesma
// Funções Recursivas necessitam de uma condição para parar
// Elas usam a pilha, caso não tenha uma condição para parar,
// Ela vai encher a pilha e vamos ter um "estouro de pilha"
package main

import (
	"fmt"
)

// Função Sequencia de Fibonnacci - 1 1 2 3 5 8 13
func fibonnacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(30)
	fmt.Println(fibonnacci(posicao))

	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonnacci(i))
	}
}
