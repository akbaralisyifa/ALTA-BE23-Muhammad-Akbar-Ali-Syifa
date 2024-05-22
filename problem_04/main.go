package main

import (
	"fmt"
)

func ubahHuruh(input string) string {

	var result string;

	for i:= 0; i < len(input); i++ {
		plusTen := input[i] + 10;
		
		// ngecek huruf kecil
		if input[i] >= 97 && input[i] <= 122 {
			// ngecek klo huruf z 
			if plusTen > 122 {
				plusTen = 97 + (plusTen - 122 - 1);
			}
			 result += string(plusTen);
		}else{
			result += string(plusTen);
		}
	}
	return result;
}

func main() {
	// text := "abc";
	// fmt.Println(text[0] + 10)
	// result := string(text[0] + 10);

	// input huruf kecil
	fmt.Println(ubahHuruh(""))
}