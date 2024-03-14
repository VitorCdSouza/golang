package main

import "fmt"

func main() {
	// sempre mesmo tipo, estatico
	var notas [3]float64
	fmt.Println(notas) // por padrao cria um array com os "zeros" do tipo fornecido

	notas[0], notas[1], notas[2] = 7.2, 5.2, 6.3
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i ++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("%.2f", media)
}