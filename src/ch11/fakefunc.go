package ch11

import (
	"fmt"
)

var realFunc = func(data string) {
	fmt.Printf("Real Func:%s\n", data)
}

func CheckInfo(data string) {
	realFunc(data)
}
