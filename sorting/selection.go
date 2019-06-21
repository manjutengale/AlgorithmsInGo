package main

import "fmt"

func selectionbubbleSort(arr *[5]int) {
	arrsize := len(arr)
	temp := 0

	for i := 0; i < arrsize-1; i++ {
		min_index := i
		for j := i + 1; j < arrsize; j++ {
			if arr[j] < arr[min_index] {
				min_index = j
			}
		}
		temp = arr[min_index]
		arr[min_index] = arr[i]
		arr[i] = temp

	}

}
func main() {
	fmt.Println("selection sort")

	arr := [5]int{5, 2, 6, 1, 9}
	fmt.Println("before sorting ", arr)
	selectionbubbleSort(&arr)
	fmt.Println("after sorting", arr)
}
