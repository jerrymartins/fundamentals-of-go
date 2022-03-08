package main

func main() {
	x, y := 1, 0
	x, y = y, x

	println(x, y)

	println(gdc(101, 10))
	println(fib(5))

}

//Greatest Common Divisor
func gdc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}
