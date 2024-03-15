package main

import "fmt"

func imprimir(aprovados ...string) {
	for _, nome := range aprovados {
		fmt.Println(nome)
	}
}

func main() {
	aprovados := []string{"Alberto", "Julia", "Fernanda"}
	imprimir(aprovados...) // tem que ter "..." para passar slice para funcoes, nao da pra passar array
}