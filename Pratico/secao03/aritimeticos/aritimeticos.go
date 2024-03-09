package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 4

	fmt.Println("Soma:", a+b)
	fmt.Println("Sub:", a-b)
	fmt.Println("Mult:", a*b)
	fmt.Println("Div:", a/b)
	fmt.Println("Modulo:", a%b)

	// bitwise
	fmt.Println("E:", a&b)
	fmt.Println("Ou:", a|b)
	fmt.Println("Xor:", a^b)

	c := 3.0
	d := 4.0

	fmt.Println("Maior:", math.Max(float64(a), float64(b))) // tem q converter para float 64 antes de passar de parametro
	fmt.Println("Menor:", math.Min(c, d))
	fmt.Println("Expodencial:", math.Pow(c, d))
}
