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
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}
}

func diaDaSemanaFall(numero int) string {
	var dia string
	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough
	case numero == 2:
		dia = "Segunda-feira"
	case numero == 3:
		dia = "Terça-feira"
	case numero == 4:
		dia = "Quarta-feira"
	case numero == 5:
		dia = "Quinta-feira"
	case numero == 6:
		dia = "Sexta-feira"
	case numero == 7:
		dia = "Sábado"
	default:
		dia = "Número Inválido"
	}

	return dia
}

func main() {
	fmt.Println("Switch")
	dia := diaDaSemanaFall(1)
	fmt.Println(dia)
}
