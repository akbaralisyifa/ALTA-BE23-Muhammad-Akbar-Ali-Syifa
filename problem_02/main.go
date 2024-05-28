package main

import "fmt"

func caesar(offside int, input string) string {

	var result string

	for _, char := range input {
		if char >= 'a' && char <= 'z' {
			shifted := char + rune(offside)

			if shifted > 'z' {
				shifted = shifted - 'z' + 'a' - 1
			}

			result += string(shifted)
		}else{
			result += string(char)
		}
	}

	return result
}

func main() {

	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))
}