package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64
type celsiusFlag struct {
	Celsius
}
type Value interface {
	String() string
	Set(string) error
}

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
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var val float64
	fmt.Sscanf(s, "%f%s", &val, &unit)
	switch unit {
	case "c", "C", "°C":
		f.Celsius = Celsius(val)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(val))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, val Celsius, usage string) *Celsius {
	f := celsiusFlag{val}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
