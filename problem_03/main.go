package main

import "fmt"

func arrayUnique(arrA, arrB []int) []int {
	var result []int;

	bMap := make(map[int]bool)
	for _, val := range arrB {
		bMap[val] = true
	}

	for _, valA := range arrA {
		if !bMap[valA] {
			result = append(result, valA)
		}
	}
	return result

}

func main() {
	fmt.Println(arrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 6, 7}))
}