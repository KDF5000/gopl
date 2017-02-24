package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Float32.Max:%g,Float64.Max:%g\n", math.MaxFloat32, math.MaxFloat64)
	fmt.Printf("%f\n", float32(16777216)+1) //16777216是1<<24，浮点数不支持位移操作，整数部分占23位，因此溢出
}
