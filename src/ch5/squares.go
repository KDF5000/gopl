package main

import (
	"fmt"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()           //f指向一个匿名函数，该匿名函数里的X是一个变量
	fmt.Println(f())         //1
	fmt.Println(f())         //4
	fmt.Println(f())         //9
	fmt.Println(squares()()) //1
	fmt.Println(squares()()) //1
}
