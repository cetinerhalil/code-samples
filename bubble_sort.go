package main

import (
	"fmt"
)

func swap(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
}

func bubbleSort(array []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array) - 1; i++ {
			if array[i + 1] < array[i] {
				swap(array, i, i + 1)
				swapped = true
			}
		}
	}	
}

func main() {

	array := []int{1, 3, 8, 4, 0, 6, 2, 9, 5, 7}
	fmt.Println("initial array ", array)
	bubbleSort(array)
	fmt.Println("Sorted array: ", array)
}