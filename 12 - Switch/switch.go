package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabádo"
	default:
		return "Número Inválido"
	}
}

func diaDaSemana2(numero int) string {
	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sabádo"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

func main() {
	fmt.Println("Switch")
	weekDay := diaDaSemana2(1)
	fmt.Println(weekDay)
	dia := diaDaSemana2(2)
	fmt.Println(dia)
	dia2 := diaDaSemana(3)
	fmt.Println(dia2)
	dia3 := diaDaSemana(4)
	fmt.Println(dia3)
	dia4 := diaDaSemana(5)
	fmt.Println(dia4)
	dia5 := diaDaSemana(6)
	fmt.Println(dia5)
	dia6 := diaDaSemana(7)
	fmt.Println(dia6)
	dia7 := diaDaSemana(1)
	fmt.Println(dia7)
	dia8 := diaDaSemana(200)
	fmt.Println(dia8)
}
