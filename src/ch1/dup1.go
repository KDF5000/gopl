package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLine(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) <= 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup1:%v\n", err)
				continue
			}
			countLine(f, counts)
			f.Close()
		}
	}
	for s, n := range counts {
		fmt.Printf("%s:%d\n", s, n)
	}
}
