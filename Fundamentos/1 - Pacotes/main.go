package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Pacote main")
	auxiliar.Escrever()

	err := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(err)
}
