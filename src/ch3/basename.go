package main

import (
	"fmt"
)

func basename(str string) string {
	var s string
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '/' {
			s = str[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
func main() {
	fmt.Println(basename("a/b/c/demo.jpg"))
}
