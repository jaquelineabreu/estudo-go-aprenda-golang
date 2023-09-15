package main

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
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Davi"
	u.idade = 21
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos devs", 255}

	//inferencia de tipo
	usuario2 := usuario{"Maria", 21, enderecoExemplo}
	fmt.Println(usuario2)

	//um campo especifico
	usuario3 := usuario{nome: "Marcia"}
	fmt.Println(usuario3)

}
