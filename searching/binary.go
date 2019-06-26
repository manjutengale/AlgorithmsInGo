package main

import "fmt"

func binary(arr [7]int, low int, high int, key int) int {
	//arrsize :=
	mid := low + (high-low)/2
	//fmt.Println("the mid element is ", mid)
	if arr[mid] == key {
		return mid
	} else if arr[mid] > key {
		return binary(arr, low, mid-1, key)
	} else {
		return binary(arr, mid+1, high, key)
	}
	return 0
}
func main() {

	arr := [7]int{1, 2, 3, 4, 5, 6, 7}
	key := 1
	fmt.Println("calling func ")
	arrsize := len(arr)
	low := 0
	high := arrsize - 1
	position := binary(arr, low, high, key)
	fmt.Println("Found the array element at the position", position)
}
