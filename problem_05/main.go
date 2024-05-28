package main

import "fmt"

func removeDuplicates(arr []int) int {

		if len(arr) == 0 {
			return 0
		}

		uniqueIndex := 1;

		for i := 1; i < len(arr); i++{

			if arr[i] != arr[uniqueIndex - 1] {
				arr[uniqueIndex] = arr[i]
				uniqueIndex++
			}
		}


		return uniqueIndex;
}

func main() {
	fmt.Println(removeDuplicates([]int{2, 3, 3, 3, 6, 9, 9}))
}