package main

import "fmt"

func mult(a, b int) int {
	return a * b
}

func exec(a func(int, int) int, p1, p2 int) int {
	return a(p1, p2)
}

func main() {
	fmt.Println(exec(mult, 10, 2))
}