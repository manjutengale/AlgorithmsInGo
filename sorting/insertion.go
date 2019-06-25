package main

import "fmt"

func insertionSort(arr *[5]int) {
	arrsize := len(arr)
	for i := 1; i < arrsize-1; i++ {
		value := arr[i]
		hole := i

		for hole > 0 && arr[hole-1] > value {
			arr[hole] = arr[hole-1]
			hole = hole - 1
		}
		arr[hole] = value
	}
}
func main() {

	arr := [5]int{5, 2, 6, 1, 9}

	insertionSort(&arr)
	fmt.Println("after sorting ", arr)
}
