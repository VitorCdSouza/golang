package main

import "fmt"

func main() {
	funcPorLetra := map[string]map[string]float64{
		"A": {
			"Ana":   1600.40,
			"André": 7895,
		},
		"B": {
			"Bernardo": 789.7,
			"Beatriz":  7845.40,
		},
	}

	for categoria, pessoas := range funcPorLetra {
		fmt.Println(categoria)
		for nome, salario := range pessoas {
			fmt.Println(nome, salario)
		}
	}
}