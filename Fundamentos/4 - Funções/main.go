package main

import "fmt"

func somar(a, b int8) int8 {
	return a + b
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	substracao := n1 - n2
	return soma, substracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, resultadoSubstracao := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma, resultadoSubstracao)
}
