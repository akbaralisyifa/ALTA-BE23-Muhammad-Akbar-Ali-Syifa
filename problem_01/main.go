package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println(a, b)
}