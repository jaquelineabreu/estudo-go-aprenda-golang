package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero int = 1000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 12345
	fmt.Println(numero3)

	//uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	//reais
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	//string
	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	//valor zero
	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	//erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
