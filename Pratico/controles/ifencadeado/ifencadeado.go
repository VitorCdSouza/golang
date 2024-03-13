package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota <= 10 && nota >= 9 {
		return "MB"
	} else if nota < 9 && nota >= 6 {
		return "B"
	} else if nota < 6 && nota >= 4 {
		return "R"
	}
	return "I"
}

func main() {
	fmt.Println(notaParaConceito(7.5))
}
