package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2

	fmt.Println(i)

	x, y := 1, 2 // := cria e inicializa
	fmt.Println(x, y)

	x, y = y, x // = somente atribui valor
	fmt.Println(x, y)
}