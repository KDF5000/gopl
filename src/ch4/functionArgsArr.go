package main

import (
	"fmt"
)

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero1(arr [32]byte) {
	for i := range arr {
		arr[i] = 1
	}
}
func main() {
	a := [32]byte{1}
	fmt.Println(a)
	zero(&a)
	fmt.Println(a)
	zero1(a)
	fmt.Println(a) //没有修改
}
