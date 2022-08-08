package main

import "fmt"

func main() {
	arr := []int{2, 14, 18, 25, 36, 41, 45, 48, 57, 68}
	target := 14
	result := binarySearch(arr, target)
	fmt.Println(result)

}

func binarySearch(arr []int, target int) int {
	start := 0
	end := len(arr) - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if target < arr[mid] {

			end = mid - 1
		} else if target > arr[mid] {

			start = mid + 1

		} else {

			return mid
		}
	}
	return -1
}
