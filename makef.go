package main

import (
	"fmt"
)

func main() {

	//slice to arr
	s := []int{1, 2, 3, 4, 5}
	var a [5]int
	copy(a[:], s)

	fmt.Println(a)

	fmt.Println("////////////////////")
	//arr to slice

	var arr = [4]int{1, 2, 3, 4}
	slice := arr[:]
	fmt.Println(slice)
}
