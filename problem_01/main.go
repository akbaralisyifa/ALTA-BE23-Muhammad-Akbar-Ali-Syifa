package main

import "fmt"

func playWithAsterik(n int) string {

	for i := 1; i <= n; i++ {
		fmt.Print("\n")
		for s := 1; s <= n - i; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i + (i - 1); j++ {
			fmt.Print("*")
		}
	}
	return ""
}

	func main(){
		playWithAsterik(5)
	}
