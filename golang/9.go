package main

import "fmt"

func linearSearch(datalist []int, key int) bool {

	for _, item := range datalist {
		if item == key {
			return true
		}
	}

	return false
}

func main() {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(linearSearch(items, 5))
}