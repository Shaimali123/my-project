package main

import "fmt"

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {

		mid := (low + high) / 2

		if haystack[mid] == needle {
			return true
		}

		if haystack[mid] < needle {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}

func main() {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(binarySearch(5, items))
	fmt.Println(binarySearch(10, items))
}