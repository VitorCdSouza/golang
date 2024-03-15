package main

import "fmt"

func fatorial(n int) (int, error) { // pode trocar int para uint, logo, nao tera retorno negativo, nao precisando retornar erro
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número Invalido %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n -1)
		return n * fatorialAnterior, nil
	}

}

func main() {
	result, _ := fatorial(5)
	fmt.Println(result)
}