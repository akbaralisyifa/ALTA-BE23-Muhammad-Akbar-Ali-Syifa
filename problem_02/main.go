package main

import "fmt";

func getMinMax(numbers ...*int)(min int, max int) {

	if len(numbers) == 0 {
		return 0, 0
	}

	min, max = *numbers[0], *numbers[0]

	for _, num := range numbers {
		if *num < min {
			min = *num
		}
		if *num > max {
			max = *num
		}
	}

	return min, max
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int

	fmt.Scanln(&a1)
	fmt.Scanln(&a2)
	fmt.Scanln(&a3)
	fmt.Scanln(&a4)
	fmt.Scanln(&a5)
	fmt.Scanln(&a6)

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6);
	fmt.Println("Nilai Minimum :" , min)
	fmt.Println("Nilai Maximum :" , max)
}