////复数
package main

import (
	"fmt"
)

func main() {
	x := 1 + 2i
	// var x complex128 = complex(1, 2) //1+2i 也可以直接写x := 1+2i
	var y complex128 = complex(2, 1) //2+1i
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
