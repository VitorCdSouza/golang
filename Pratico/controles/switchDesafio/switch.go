package main

import "fmt"

func notaParaConceito(nota float64) string {
	switch {
	case nota <= 10 && nota >= 9:
		return "MB"
	case nota < 9 && nota >= 6:
		return "B"
	case nota < 6 && nota >= 4:
		return "R"
	case nota < 4 && nota >= 0:
		return "I"
	default:
		return "Erro"
	}
}

func main() {
	fmt.Println(notaParaConceito(7))
}