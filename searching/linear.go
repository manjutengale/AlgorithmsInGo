package main

import "fmt"

func linear(arr [5]int, key int) {
	fmt.Println("the key of element to be found is ", key)
	arrsize := len(arr)

	for i := 0; i <= arrsize-1; i++ {
		if arr[i] == key {
			fmt.Println("key found at postition ", i)
			break
		}
	}

}
func main() {
	arr := [5]int{5, 2, 6, 1, 9}
	key := 9
	linear(arr, key)
}
