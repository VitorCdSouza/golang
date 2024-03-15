package main

import "fmt"

func main() {
	funciosSalarios := map[string]float64{
		"Jo":      11564.10,
		"Matheus": 1000,
		"Maria":   8792.10,
	}
	fmt.Println(funciosSalarios)

	funciosSalarios["Gabriel"] = 3000
	delete(funciosSalarios, "ine")

	for nome, salario := range funciosSalarios {
		fmt.Println(nome, salario)
	}
}
