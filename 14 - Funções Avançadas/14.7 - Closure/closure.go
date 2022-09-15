// São Funções que referenciam váriaveis que estão fora de seu corpo
package main

import (
	"fmt"
)

func closure() func() {
	texto := "dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcNova := closure()
	funcNova()
}
