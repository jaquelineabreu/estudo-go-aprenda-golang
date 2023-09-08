package main

import "fmt"

func fibonacci(posicao uint) int {
	if posicao <= 1 {
		return int(posicao)
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(12)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
