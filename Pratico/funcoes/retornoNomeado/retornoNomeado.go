package main

import "fmt"

func trocar(p1, p2 int) (seg int, pri int) {
	pri = p1
	seg = p2
	return
}

func main() {
	var1, var2 := trocar(1, 2)
	fmt.Println(var1, var2)
}