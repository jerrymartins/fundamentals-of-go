package main

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FeezingC      Celsius = 0
	BoilingC      Celsius = 100
)

func main() {
	f := CtoF(500)
	c := FToC(32)

	println(f)
	println(c)
}

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
