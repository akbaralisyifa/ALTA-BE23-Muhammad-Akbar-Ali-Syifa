package main

import "fmt"

func main() {

	var bilangan int;
	fmt.Println("Input bilangan: ")
	fmt.Scanln(&bilangan)
	fmt.Println("~ hasil 2.1 ~")

	// problem 2.1 => bilangan mengurut ke atas
	for i := 1; i <= bilangan; i++ {
		if (bilangan % i == 0) {
			fmt.Println(i)
		}
	}

	fmt.Println("~ hasil 2.2 ~")
	// problem 2.2 => bilangan mengurut ke bawah
	for i := bilangan; i >= 1; i--{
		if(bilangan % i == 0){
			fmt.Println(i)
		}
	}
}