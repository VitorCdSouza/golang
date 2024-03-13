package main

import "fmt"

// variavel é uma referencia para um local de memória
// ponteiro aponta para o local de memória que é referenciado em variaveis

func main() {
	i := 1 // inteira, ocupa 64 bits / 8 bytes

	var p *int = nil // nil = null

	p = &i 
	fmt.Println(*p) // * acessa o valor do endereço da memória

	*p++
	fmt.Println(i)
	fmt.Println(p, &i)
	// Go n tem aritimetica de ponteiro
	// ponteiro pode apontar para enderecos nulos
}