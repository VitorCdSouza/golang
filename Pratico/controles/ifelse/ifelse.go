package main

import "fmt"

func imprimirResult(nota float64) {
	if nota >= 7 { // if não precisa de parenteses, tem que ter {}
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Reprovado")
	}
}

func main() {
	imprimirResult(7)
}