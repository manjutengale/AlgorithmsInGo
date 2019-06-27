package main

import (
	"fmt"
	"math"
)

func jump(arr []int, key int) int {
	var sizeinInt int
	sizeinInt = len(arr)
	var sizeinFloat float64
	sizeinFloat = float64(sizeinInt)
	jumpsize := math.Sqrt(sizeinFloat)
	leftside := 0
	rightside := int(jumpsize)
	fmt.Println("the jumpsize is ", rightside)

	for arr[rightside] <= key && rightside < sizeinInt {
		leftside = rightside
		rightside += int(math.Sqrt(sizeinFloat))
		fmt.Println(rightside)
		if rightside > sizeinInt-1 {
			rightside = sizeinInt
		}
	}
	for i := leftside; i < rightside; i++ {
		if arr[i] == key {
			return i
		}
	}

	return -1
}
func main() {

	arr := []int{1, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610}
	key := 2
	positionjump := jump(arr, key)
	fmt.Println("found the element at the index ", positionjump)
}
