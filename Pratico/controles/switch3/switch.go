package main

import "fmt"

func tipoVar(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "n sei"
	}
}

func main() {
	fmt.Println(tipoVar(312))
	fmt.Println(tipoVar(func() {}))
	fmt.Println(tipoVar("Oi"))
}