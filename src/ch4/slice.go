package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 3, 4, 5}
	a2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(cap(a1)) //5
	fmt.Println(cap(a2)) //5
	b1 := []int{}
	b2 := [...]int{}
	fmt.Println(cap(b1)) //0
	fmt.Println(cap(b2)) //0
}
