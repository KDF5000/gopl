package main

import (
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (i *IntSet) Add(x int) {
	word, bits := x/64, uint(x%64)
	for word >= len(i.words) {
		i.words = append(i.words, 0)
	}
	i.words[word] |= (1 << (63 - bits))
}

func (i *IntSet) Delete(x int) {
	word, bits := x/64, uint(x%64)
	if word < len(i.words) {
		i.words[word] &^= (1 << (63 - bits))
	}
}

func (i *IntSet) Has(x int) bool {
	word, bits := x/64, uint(x%64)
	return (word < len(i.words) && i.words[word]&(1<<(63-bits)) != 0)
}

func (i *IntSet) UnionSet(s *IntSet) {
	for word, _ := range s.words {
		if word >= len(i.words) {
			i.words = append(i.words, s.words[word])
		} else {
			i.words[word] |= s.words[word]
		}
	}
}

func (i *IntSet) ToString() string {
	var s string
	for _, v := range i.words {
		s += fmt.Sprintf("%b,", v)
	}
	return s
}

func (i *IntSet) Set() []int {
	nums := make([]int, 0)
	for k, word := range i.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if i.words[k]&(1<<(63-uint(j))) != 0 {
				nums = append(nums, 64*k+j)
			}
		}
	}
	return nums
}

func main() {
	intSet := &IntSet{}
	intSet.Add(0)
	intSet.Add(2)
	intSet.Add(64)
	intSet.Add(63)
	intSet.Add(100)

	fmt.Println(intSet.Has(1))
	fmt.Println(intSet.Has(2))
	fmt.Println(intSet.Has(0))
	fmt.Println(intSet.ToString())
	intSet.Delete(2)
	fmt.Println(intSet.ToString())

	intSet2 := &IntSet{}
	intSet2.Add(1)
	intSet2.Add(0)
	intSet2.Add(130)
	fmt.Println(intSet2.ToString())

	intSet.UnionSet(intSet2)
	fmt.Println(intSet.ToString())

	fmt.Println(intSet.Set())
}
