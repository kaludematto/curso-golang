package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)
	// // for init
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// RANGE
	nomes := [3]string{"Kalu", "Rosi", "Irene"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	// Iterar por Strings
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra) // Retorna o numero da tabela ASCII
	}

	// iterar por um map
	usuario := map[string]string{
		"nome":      "Kalu",
		"sobrenome": "Dematto",
	}
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	// Não é possivel usar range em structsE
}
