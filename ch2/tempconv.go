// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点
	BoilingC      Celsius = 100     // 沸点
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func main() {
	c := FToC(212.0)
	fmt.Println(c)
}
