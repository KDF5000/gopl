package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZrorC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //结冰点温度
	BoilingC      Celsius = 100     //沸点
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g F", f)
}
