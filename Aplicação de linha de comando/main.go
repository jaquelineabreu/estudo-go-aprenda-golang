package main

//EX: comandos
//go run main.go ip --host amazon.com.br
//go run main.go servidores --host amazon.com.br

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
