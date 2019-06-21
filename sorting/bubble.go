package main

import "fmt"

func bubbleSort(arr *[5]int) {
	arrsize := len(arr)
	temp := 0
	for i := 0; i < arrsize-1; i++ {
		for j := 0; j < arrsize-i-1; j++ {
			if arr[j] > arr[j+1] {
				temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	println(arrsize)
}
func main() {
	fmt.Println("bubble sort")

	arr := [5]int{5, 2, 6, 1, 9}
	fmt.Println("before sorting ", arr)
	bubbleSort(&arr)
	fmt.Println("after sorting", arr)
}
