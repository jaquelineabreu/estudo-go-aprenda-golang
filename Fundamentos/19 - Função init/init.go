package main

import "fmt"

var n int

func init() {
	fmt.Println("Init")
	n = 10
}

func main() {
	fmt.Println("Main")
	fmt.Println(n)
}
