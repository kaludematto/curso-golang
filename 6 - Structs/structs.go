package main

// Struct é uma coleção de campos em Go
// Similar as Classes em outras linguagens
// Usado quando é necessário salvar mais de um valor

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Hellow")

	var u usuario
	u.nome = "Kalu"
	u.idade = 35
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Armando Ramos Filho", 177}

	u2 := usuario{"Kalu", 35, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Kalu"}
	fmt.Println(u3)

	u4 := usuario{idade: 35}
	fmt.Println(u4)
}
