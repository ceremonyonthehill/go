package main

import (
	"fmt"
	"strings"
)

func main() {

	var qa [5]int = [5]int{10, 20, 30, 40, 50}

	qs := qa[:]

	var nq = make([]int, 0, 10)

	nq = append(nq, qs...)

	var fixedArray [5]int
	copy(fixedArray[:], nq)

	inventory := make(map[string]int)

	itemNames := []string{"john", "paul", "george", "ringo", "billy"}
	for i, beatles := range itemNames {
		lowerName := strings.ToLower(beatles)
		inventory[lowerName] = fixedArray[i]
	}

	checkItems := []string{"john", "paul", "ringo"}
	for _, item := range checkItems {
		qty, ok := inventory[item]
		if ok {
			fmt.Printf("%s: %d units\n", item, qty)
		} else {
			fmt.Printf("%s: not found in inventory\n", item)
		}
	}

	fmt.Println("Capacity before append:", cap(nq))
	for i := 0; i < 10; i++ {
		nq = append(nq, 1)

	}
	fmt.Println("Capacity after append:", cap(nq))

}
