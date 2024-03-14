package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // tamanho fixo
	s1 := []int{1, 2, 3}  // tamanho variavel
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// slice != array, slice é um pedaço de um array

	s2 := a2[1:3] // slice sempre aponta para um mesmo array, é um ponteiro
	fmt.Println(s2)

	s3 := a2[0:2]
	fmt.Println(s3)

	s4 := s3[:1]
	fmt.Println(s4)
}