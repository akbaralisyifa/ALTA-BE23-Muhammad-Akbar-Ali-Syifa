package main

import (
	"fmt"
	"strconv"
)

// fungsi untuk cek bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	};

	for i:= 2; i < n; i++{
		if n % i == 0 {
			return false;
		}
	}

	return true
}

// fungsi untuk mengecek bilangan fullprima
func isFullPrime(n int) bool {
	if !isPrime(n){
		return false
	};

	strConv := strconv.Itoa(n);  // slice, array, string ( yang bisa di mapping )
		for _, char := range strConv {
			nilai, err := strconv.Atoi(string(char));
			
			if err != nil {
				fmt.Println(err.Error()) // pesan bawaan dari error default
				return false;
			}

		if !isPrime(nilai) {
			return false
		}
	}
	return true;
}

func main() {
	// Test cases
	numbers := []int{2, 3, 11, 13, 23,29, 37, 41, 43, 53};

	for _, num := range numbers { // for index, nilai = ...
		fmt.Println(isFullPrime(num));
	}
}