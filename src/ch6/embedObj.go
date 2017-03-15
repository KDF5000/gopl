package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Color struct {
	R, G, B int
}

//三维点
type DPoint struct {
	X, Y, Z float64
}

type ColoredPoint struct {
	Point //只有匿名才才运行嵌入该类型的对象向访问自己成员一样访问其Point成员
	Color Color
	DPoint
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//测试同名函数，
// func (c DPoint) Distance(m DPoint) float64 {
// 	return math.Sqrt(math.Pow(c.X, m.X) + math.Pow(c.Y, m.Y) + math.Pow(c.Z, m.Z))
// }

func main() {
	cp := ColoredPoint{Point{0, 0}, Color{0, 0, 0}, DPoint{0, 0, 0}}
	cq := ColoredPoint{Point{2, 2}, Color{0, 0, 0}, DPoint{1, 1, 1}}
	fmt.Println(cp.Distance(cq.Point)) //编译错误，ambiguous selector cp.Distance

	p1 := Point{0, 0}
	p2 := Point{1, 2}
	fmt.Println(math.Hypot(p2.X-p1.X, p2.Y-p1.Y))
}
