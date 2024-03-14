package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta os elementos inseridos na inicializacao e gera um array com tais valores

	for i, numero := range numeros { // range retorna o indice e o elemento
		fmt.Printf("%d-> %d\n", i, numero)
	}

	for _, numero := range numeros { // "_" para ignorar o retorno
		fmt.Printf("%d\n", numero)
	}
}