package main

import "fmt"

func obterAprovado(n int) bool {
	defer fmt.Println("Fim")
	if n < 10 {
		fmt.Println("oi")
		return true
	}
	return false
}

func main() {
	fmt.Println(obterAprovado(4))
}