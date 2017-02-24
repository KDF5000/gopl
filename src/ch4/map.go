package main

import (
	"fmt"
	"sort"
)

func main() {
	args := map[string]int{
		"Tom":   21,
		"Alice": 12,
		"Bob":   14,
	}
	names := make([]string, 0, len(args))
	for name := range args {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s:%d\n", name, args[name])
	}
	//
	var ages map[string]int //nil
	fmt.Println(len(ages))  //0
	// ages["Tom"] = 32        //panic: assignment to entry in nil map
	ages = make(map[string]int)
	ages["Tom"] = 32 //Ok
	age, ok := ages["bob"]
	if !ok {
		fmt.Println("No key:  Bob,", age)
	}
}
