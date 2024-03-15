package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[321456789] = "Jao"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Println("Nome:", nome, " Cpf:", cpf)
	}

	delete(aprovados, 321456789)
	fmt.Println(aprovados)
}