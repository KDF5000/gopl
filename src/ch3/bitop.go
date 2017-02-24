package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 &^ 0) //1
	fmt.Println(1 &^ 1) //0

	///
	test()
}
func test() {
	var x uint8 = 1<<1 | 1<<5 //00100010
	var y uint8 = 1<<1 | 1<<2 //00000110
	fmt.Printf("%08b\n", x)   // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)   // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)   // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)   // "00100110", the union {1, 2, 5} //或
	fmt.Printf("%08b\n", x^y)   // "00100100", the symmetric difference {2, 5}不同为1
	fmt.Printf("%08b\n", x&^y)  // "00100000", the difference {5} 如果y对应的位为1，则为0；否则为x的值
	fmt.Printf("%08b\n", -5>>1) //  -0000011

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}
