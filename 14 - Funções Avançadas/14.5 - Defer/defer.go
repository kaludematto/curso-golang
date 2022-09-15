package main

import "fmt"

func funcao1() {
	fmt.Println("Chamando a Função 1")
}

func funcao2() {
	fmt.Println("Chamando a Função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada. Resultado será retornado")
	fmt.Println("Entrando na Função para ver se o Aluno está aprovado!")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// fmt.Println("Defer")
	// // DEFER = ADIAR
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
