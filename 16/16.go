package main

import "fmt"

func main() {
	unsorted := []int{63, 5, 1, 7, 10, 15, 6}
	fmt.Println(unsorted)
	sorted := quickSort(unsorted)
	fmt.Println(sorted)
}

func quickSort(array []int) []int {
	for i := 0; i < len(array); i += 1 {
		for j := 0; j < len(array); j += 1 {
			if array[i] > array[j] {
				continue
			} else {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}
