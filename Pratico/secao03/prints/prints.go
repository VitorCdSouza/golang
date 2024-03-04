package main

import "fmt"

func main() {
	fmt.Print("Mesma") 
	fmt.Print(" Linha") // "Print", printa na mesma linha

	fmt.Println(" Nova.") // "Println" printa e quebra a linha
	fmt.Println("Linha")

	x := 45646.32123
	xs := fmt.Sprint(x) // transforma em string, tem versão Sprintf
	fmt.Println("O valor de x é: " + xs) // concatenacao com "+" só funciona para string
	fmt.Println("O valor é:", x) // concatencao com "," ja da o espaco automatico
	
	fmt.Printf("O valor de x é: %.2f", x)

	a := 123
	b := 123.231
	c := true
	d := "Oi"
	fmt.Printf("\na: %d, b: %f, b: %.2f, c: %t, d: %s", a, b, b, c, d)
}