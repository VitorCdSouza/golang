package main

import (
	"fmt"
	"time"
)

func main() {
	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"aa"}
	p2 := Pessoa{"aa"}

	fmt.Println(p1 == p2)
	fmt.Println(&p1)
}