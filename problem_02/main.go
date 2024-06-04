package main

import "fmt"

func moneyChange(money int) []int {

	denominations := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	var result []int

	for _, denom := range denominations {
		count := money / denom

		for i := 0; i < count; i++ {
			result = append(result, denom)
		}

		money -= count * denom
	}

	return result

}

func main() {
	fmt.Println(moneyChange(123))
	fmt.Println(moneyChange(432))
	fmt.Println(moneyChange(543))
	fmt.Println(moneyChange(7752))
	fmt.Println(moneyChange(15321))
}