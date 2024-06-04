package main

import "fmt"

func sampleEquations(a, b, c int) {
	found := false

	for x:= -100; x <= 100; x++{
		for y:= -100; y <= 100; y++{
			for z:= -100; z <= 100; z++{
				if x+y+z == a && x*y*z == b && x*x+y*y+z*z == c{
					fmt.Printf("Output: %d %d %d\n", x, y, z)
					found = true
					return
				}
			}
		}
	}

	if !found {
		fmt.Println("no solution")
	}
}

func main() {
	sampleEquations(1, 2, 3) // no solution
	sampleEquations(6, 6, 14)
}