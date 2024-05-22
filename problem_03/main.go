package main

import "fmt"

func cetakTabelPerkalian(n int) {

	for i := 1; i <= n; i++ {
		for j:=1; j <= n; j++ {
			fmt.Print(i * j, "\t" )
		}
		fmt.Print("\n")
	}
}

func main() {
	cetakTabelPerkalian(9)
}