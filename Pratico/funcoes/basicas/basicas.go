package main

import "fmt"

func f1() {
	fmt.Println("Oi")
}

func f2(p1 string) {
	fmt.Println(p1)
}

func f3() string {
	return "oi"
}

func f4(t1, t2 string) string {
	return t1 + " " + t2
}

func f5() (string, string) {
	return "1", "2"
}

func main() {
	f1()
	f2("aaa")
	fmt.Println(f3())
	fmt.Println(f4("oi", "oi"))
	fmt.Println(f5())
}