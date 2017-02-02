package main

import "fmt"
    
func swap(array []int, i, j int) {
	temp := array[j]
	array[j] = array[i]
	array[i] = temp
}

func selectionSort(array []int) {
	
	for i := 0; i <= len(array) - 1; i++ {
		min := i
		for j := i + 1; j <= len(array) - 1; j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		swap(array, i, min)
	}
}

func main() {

	array := []int{1, 3, 8, 4, 0, 6, 2, 9, 5, 7}
	fmt.Println("Initial array: ", array)
	selectionSort(array)
	fmt.Println("Sorted array: ", array)
}
