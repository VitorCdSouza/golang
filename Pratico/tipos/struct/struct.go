package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 { // o (p produto) antes do nome da funcao declara a quem a funcao ira pertencer
	return p.preco * (1 - p.desconto)
}

func main() {
	var p1 produto = produto{
		"a",
		7895.4,
		0.15,
	}

	fmt.Println(p1.precoComDesconto())

	p2 := produto{
		"b",
		123.23,
		0.10,
	}

	fmt.Println(p2.precoComDesconto())
}
