package main

import (
	"fmt"
	m "math" // nome antes do import cria apelido
	// "_" serve para nao remover o import na hora da compilacao do import
)

func main() {
	const PI float64 = 3.1415 // const <nome> <tipo> = <valor>
	var raio = 3.2            // nao precisa explicitar o tipo da variavel

	area := PI * m.Pow(raio, 2) // obrigatorio o uso da variavel criada em go
	fmt.Println("A area é:", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	g, h := 1, true
	fmt.Println(g, h)
}
