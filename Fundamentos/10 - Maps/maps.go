package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	//para acessar a chave
	fmt.Println(usuario["nome"])

	//map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"Campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	//deletar chave
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	//adicionar uma chave
	usuario2["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)

}
