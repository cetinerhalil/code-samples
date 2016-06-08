package main

import "fmt"

func swap(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
}

func shellSort(array []int) {
	x := 1
	for x < len(array) {
		x = 3 * x + 1
	}
	for x >= 1 {
		for i := x; i < len(array); i++ {
			for j := i; j >= x && array[j] < array[j - x]; j = j - x {
				swap(array, j, j - x)
			}
		}
		x = x/3
	}
}

func main() {
  array := []int{1, 3, 8, 4, 0, 6, 2, 9, 5, 7}
  fmt.Println("Initial array: ", array)
  shellSort(array)
  fmt.Println("Sorted array: ", array)
}