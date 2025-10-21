package main

import (
	"fmt"
)

func main() {
	pam := make(map[int]bool)
	pam[0] = true
	v, ok := pam[0]
	if ok {
		fmt.Println("Value:", v)
	}
	delete(pam, 0)

}
