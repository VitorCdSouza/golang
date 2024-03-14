package main

import "fmt"

func main() {
	s1 := make([]int, 10) // cria inteirnamente um array e esse slice aponta para ele
	s1[9] = 12
	fmt.Println(s1)

	s2 := make([]int, 10, 20) // faz a mesma coisa, porem cria um array maior
	fmt.Println(s2, len(s2), cap(s2)) // cap() exibe a capacidade do array referenciado pelo slice

	s2 = append(s2, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(s2, len(s2), cap(s2))

	s2 = append(s2, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9) // quando slice atinge limite dos elementos do array, o go cria outro array e referencia com o mesmo slice
	fmt.Println(s2, len(s2), cap(s2))
}