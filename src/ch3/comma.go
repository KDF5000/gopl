package main

import (
	"fmt"
	"strings"
)

// 1234 -> 1,234   1234.5434 -> 1,234.5434
func comma(num string) string {
	n := len(num)
	if n <= 3 {
		return num
	}
	dot := strings.LastIndex(num, ".")
	// fmt.Printf("%d,%T\n", dot, dot)
	if dot == -1 {
		return comma(num[:n-3]) + "," + num[n-3:]
	} else {
		return comma(num[:dot]) + num[dot:]
	}
	return num
}
func main() {
	fmt.Println(comma("12345.3224"))
}
