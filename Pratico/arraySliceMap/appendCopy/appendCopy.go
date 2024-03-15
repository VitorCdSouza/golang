package main

import "fmt"

func main() {
	a1 := [3]int{1, 2, 3}
	fmt.Println(a1)
	var s1 []int

	s1 = append(s1, 4, 5, 6)
	fmt.Println(s1)

	s2 := make([]int, 2) // slice com 2 elementos copiou só 2 elementos do slice "original"
	copy(s2, s1) // copy não aumenta o slice
	fmt.Println(s2)
}