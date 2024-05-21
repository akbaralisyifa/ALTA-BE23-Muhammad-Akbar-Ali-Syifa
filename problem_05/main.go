package main

import (
	"fmt"
	"math"
)
  
func pangkat(x, n int) int {
	if n == 0 {
		return x;
	}
	return int(math.Pow(float64(x), float64(n)));
}

func main(){
	fmt.Println(pangkat(2, 3));
	fmt.Println(pangkat(5, 3));
	fmt.Println(pangkat(10, 2));
	fmt.Println(pangkat(2, 5));
	fmt.Println(pangkat(7, 3));
}