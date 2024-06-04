package main

import "fmt"

func BinarySearch(array []int, x int) int {
	begin := 0
	end := len(array) - 1

	for begin <= end {
		mid := begin + (end-begin)/2
		if array[mid] == x {
			return mid
		} else if array[mid] < x {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	fmt.Println(BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))
	fmt.Println(BinarySearch([]int{12, 15, 19, 24, 31, 53, 59, 60}, 31))
	fmt.Println(BinarySearch([]int{12, 15, 19, 24, 31, 53, 59, 60}, 100))
}