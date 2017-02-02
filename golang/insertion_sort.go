package main

import "fmt"

func swap(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j - 1]; j-- {
			swap(array, j, j - 1)
		}
	}
}

func main() {
  
  array := []int{1, 3, 8, 4, 0, 6, 2, 9, 5, 7}
  fmt.Println("Initial array: ", array)
  insertionSort(array)
  fmt.Println("Sorted array: ", array)
}
