package main

import "fmt"

func media(numero ...float64) (media float64) {
	total := 0.0
	for _, num := range numero {
		total += num
	}
	media = total / float64(len(numero))
	return
}

func main() {
	fmt.Println(media(10, 2, 7))
}