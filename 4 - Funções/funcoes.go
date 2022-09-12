package main

import "fmt"

func main() {
	soma := somar(19, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Ola")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(30, 20)
	fmt.Println(resultadoSoma, resultadoSubtracao)
	resultadoSoma2, _ := calculosMatematicos(30, 20)
	fmt.Println(resultadoSoma2)
}

// Função que retorna mais de um valor
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}
