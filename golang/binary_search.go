package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(list ...int, toFind int) int {

	var low, high int
	low = list[0]
	high = list[len(list)-1]

	if toFind > high || toFind < low {
		return -1
	}

	for low <= high {
		mid := low + (high-low)/2
		switch {
		case toFind < list[mid]:
			high = mid - 1

		case toFind > list[mid]:
			low = mid + 1

		case toFind == list[mid]:
			return mid
		}
	}
	return -1
}

func makeSlice(a int, b int) []int {
	var list []int
	for n := a; n <= b; n++ {
		list = append(list, n)
	}
	return list
}


func main() {
	fmt.Println("Binary search:")
	list := makeSlice(1, 99)
	
	rand.Seed(time.Now().UTC().UnixNano())
	toFind := rand.Intn(100)

	index := binarySearch(list, toFind)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("list[", index, "] = ", list[index])
	}
}
