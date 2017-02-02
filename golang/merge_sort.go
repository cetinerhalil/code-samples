package main

import "fmt"

func mergeSort(a []int) []int {

  if len(a) <= 1 {
    return a
  }
  
  left := make([]int, 0)
  right := make([]int, 0)
  m := len(a) / 2

  for i, x := range a {
    switch {
    case i < m:
      left = append(left, x)
    case i >= m:
      right = append(right, x)
    }
  }

  left = mergeSort(left)
  right = mergeSort(right)

  return merge(left, right)
}

func merge(left, right []int) []int {

  final := make([]int, 0)

  for len(left) > 0 || len(right) > 0 {
    if len(left) > 0 && len(right) > 0 {
      if left[0] <= right[0] {
        final = append(final, left[0])
        left = left[1:len(left)]
      } else {
        final = append(final, right[0])
        right = right[1:len(right)]
      }
    } else if len(left) > 0 {
      final = append(final, left[0])
      left = left[1:len(left)]
    } else if len(right) > 0 {
      final = append(final, right[0])
      right = right[1:len(right)]
    }
  }

  return final
}

func main() {
  
  array := []int{1, 3, 8, 4, 0, 6, 2, 9, 5, 7}
  fmt.Println("Initial array: ", array)
  array = mergeSort(array)
  fmt.Println("Sorted array: ", array)
}
