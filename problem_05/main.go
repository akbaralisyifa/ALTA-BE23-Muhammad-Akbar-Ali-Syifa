package main

import "fmt"

// menghitung luas lingkaran

func main(){
	var (
		T float64 = 20;
		r float64 = 4;
	 	pi float64 = 3.14;
	)

	Lp := 2 * pi * r * (r+T);

	fmt.Println("Luas lingkaran =", Lp)
}