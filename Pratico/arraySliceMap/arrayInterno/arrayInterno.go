package main

import "fmt"

func main() { 
	s1 := make([]int, 10, 15)
	s2 := append(s1, 1, 2, 3)
	fmt.Println(s1, s2)

	s1[3] = 4
	fmt.Println(s1, s2)
}