package main

import "fmt"

func notaParaConceito(n float64) string {
	nota := int(n)
	switch nota {
	case 10:
		fallthrough // tem que ter fallthrough para passar ao proximo case, se nao, ira agir como um break por padrao
	case 9:
		return "MB"
	case 8, 7, 6:
		return "B"
	case 5, 4:
		return "R"
	case 3, 2, 1, 0:
		return "I"
	default:
		return "Nota Invalida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9))
}