package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) { // nao recomendado, nao é funcao pura
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	inc2(&n)
	fmt.Println(n)

	// inc2 altera o valor no endereco da memoria, inc1 apenas no escopo da funcao
}