package main

import "fmt"

func iterativebinary(arr []int, key int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] == key {
			return mid
		}
		if arr[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}
	return -1
}
func binary(arr []int, low int, high int, key int) int {
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
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	key := 5
	fmt.Println("calling func ")
	arrsize := len(arr)
	low := 0
	high := arrsize - 1
	positioniterative := iterativebinary(arr, key)
	position := binary(arr, low, high, key)
	fmt.Println("Found the array element at the position", position)
	fmt.Println("iter Found the array element at the position", positioniterative)
}
