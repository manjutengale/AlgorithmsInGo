package main

import "fmt"

func binary(arr *[7]int, key int, low int, high int) int {
	//arrsize :=
	mid := low + high/2
	fmt.Println("the mid element is ", mid)
	if arr[mid] == key {
		return mid
	} else if arr[mid] > key {
		return binary(arr, key, low, mid-1)
	} else {
		return binary(arr, key, mid+1, high)
	}
	return 0
}
func main() {

	arr := [7]int{15, 2, 54, 1, 7, 22, 8}
	key := 54
	fmt.Println("calling func ")
	arrsize := len(arr)
	low := 0
	high := arrsize - 1
	position := binary(&arr, key, low, high)
	fmt.Println("Found the array element at the position", position)
}
