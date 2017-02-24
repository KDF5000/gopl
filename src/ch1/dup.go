package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //ctrl+d结束
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Println(n, ":", line)
		}
	}
}
