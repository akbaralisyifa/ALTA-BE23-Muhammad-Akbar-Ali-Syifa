package main

import "fmt"

func drawXYZ(n int) string {

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			position := (i-1)*n + j
			if position%2 == 0 {
				fmt.Print("Z")
			} else if position%3 == 0 {
				fmt.Print("Y")
			} else {
				fmt.Print("X")
			}
		}

		fmt.Print("\n")
	}

	return ""
}

func main() {
	drawXYZ(5)
}