package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("String: " + string(1232)) // esse string converte código da tabela ASCII

	fmt.Println("String correta " + strconv.Itoa(123)) // da pra converter com o fmt ou concatenando com ","

	// string pra int
	num, erro := strconv.Atoi("123") // dois valores, o valor e um erro
	fmt.Println(num - 122, erro)

	// strconv.ParseBool, converte para booleano
}