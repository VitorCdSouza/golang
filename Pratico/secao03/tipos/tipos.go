package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 10000)
	fmt.Println("Literal inteiro, tipo: ", reflect.TypeOf(32000))

	// naturais inteiros positivos: uint8 uint16 uint32 uint64

	// byte em golang é um uint8
	
	// valor máximo de um int em go é: 9223372036854775807

	// tipo "rune" é um int32, altera um string para código da tabela Unicode

	// por padrao float é float64

	// strings são limitadas apenas por aspas duplas
	// metodo len() da o tamanho de uma string
	// string com mutiplas linhas:
	s2 := `a
	b
	c
	d
	`

	fmt.Println(s2)

	// nao há tipo char em go
}