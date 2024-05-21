package main

import (
	"fmt"
	"unicode/utf8"
)

func palindrom(input string) bool {
	rune := utf8.RuneCountInString(input);
	middle := rune / 2;

	for i := 0; i < middle; i++ {
		if input[i] != input[rune - 1 - i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrom("civic"))
	fmt.Println(palindrom("katak"))
	fmt.Println(palindrom("kasur rusak"))
	fmt.Println(palindrom("kupu-kupu"))
	fmt.Println(palindrom("lion"))

}