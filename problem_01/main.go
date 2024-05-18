package main

import "fmt"

func main() {
	var name string
	var nilai int

	fmt.Println("Input your name: ")
	fmt.Scanln(&name)
	fmt.Println("Input your nilai: ")
	fmt.Scanln(&nilai)

	switch {
		case (nilai >= 80):
			fmt.Println("nama:", name, ", Nilai: A")
		case (nilai >= 65):
			fmt.Println("nama:", name, ", Nilai: B+")
		case(nilai >= 50):
			fmt.Println("nama:", name, ", Nilai: B")
		case(nilai >= 35):
			fmt.Println("nama:", name, ", Nilai: C")
		default:
			fmt.Println("nama:", name, ", Nilai: D")
	}
}