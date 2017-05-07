package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//生成0-99的自然数
	//counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	//计算自然数的平方
	go func() {
		//只要naturals有数据就一直循环，知道close掉
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
