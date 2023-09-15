package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	//formas de declarar array
	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 0", "Posição 1", "Posição 2", "Posição 3", "Posição 4"}
	fmt.Println(array2)

	//forma alternativa
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//formas de declarar slice
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	//slice apontando para posições do array
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//Arrays Internos
	//tipo-tamanho-capacidade
	fmt.Println("---------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
