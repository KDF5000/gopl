package main

import (
	"fmt"
)

func findComplement(num int) int {

	return 0
}

func main() {
	fmt.Printf("%b\n", ^(-4))                     //binary: 11
	fmt.Printf("%b\n", ^1)                        //binary: -10
	fmt.Printf("%b\n", (-1)^1)                    //binary: -10
	fmt.Printf("%b\n", -5^5)                      //binary: -10
	fmt.Printf("%b\n", 5&0xffffffff)              //binary: 101
	fmt.Printf("%T,%T\n", 0xf, 0xfffffffffffffff) //int, int
	var a int8
	a = -5
	fmt.Printf("%b\n", a^(0xf))     //binary: -1100
	fmt.Printf("%b\n", a^(-0xf))    //binary: 1010
	fmt.Printf("%b,%b\n", -0xf, -5) //binary: 1011
	fmt.Printf("%b\n", ^5)          //binary: 1011
	// fmt.Printf("%b\n", a^(0xff)) //constant 255 overflows int8
	// fmt.Printf("%b\n", ^-1)        //0
	// fmt.Println(findComplement(5)) //101 --> 010 / 2
	var b int
	b = 5
	fmt.Printf("%b,%b\n", -(^b), b)
	fmt.Printf("%b\n", (-(^b))&b)
}
