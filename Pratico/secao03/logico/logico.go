package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv := trab1 && trab2
	comprarTv32 := trab1 != trab2 // XOR ou exclusico
	comprarSorvete := trab1 || trab2

	return comprarTv, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Println(tv50, tv32, sorvete)
}