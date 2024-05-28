package main

import "fmt"

func findMaxSubArray(k int, arr []int) int {
	n := len(arr)

	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	windowSum := maxSum
	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum

}

func main() {
	fmt.Println(findMaxSubArray(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println(findMaxSubArray(2, []int{2, 3, 4, 1, 5}))
}