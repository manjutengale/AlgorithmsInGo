package main

import "fmt"

func isPalindrome(x int) bool {
	fmt.Println("the num is ", x)

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertednumber := 0
	for x > revertednumber {
		revertednumber = revertednumber*10 + x%10
		x = x / 10

	}
	fmt.Println("the revertednumber is ", revertednumber)
	fmt.Println("the revertednumber/10 is ", revertednumber/10)
	return x == revertednumber || x == revertednumber/10
	// if x == revertednumber || x == revertednumber/10 {
	// 	return true
	// } else {
	// 	return false
	// }
}
func main() {

	num := 10
	isPal := isPalindrome(num)
	fmt.Println("the valus is ", isPal)
}
