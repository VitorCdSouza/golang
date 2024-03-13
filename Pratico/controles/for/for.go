package main

import (
	"fmt"
	"time"
)

func main() {
	// em go o unico laço de repeticao é o for

	i := 1
	for i <= 10 {
		fmt.Println(i)
		i ++
	}

	for a := 0; a <= 20; a ++ { 
		fmt.Println(a)
	}

	// evitar inserir estruturas umas dentro das outras
	for i := 0; i < 10; i ++ {
		if i % 2 != 0 {
			fmt.Println(i, "impar")
		}
	}

	for {
		// laco infinito
		fmt.Println("Infinito")
		time.Sleep(time.Second) // repete de um em um segundo
	}
}