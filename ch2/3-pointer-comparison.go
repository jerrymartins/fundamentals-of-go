package main

func main() {
	var x, y int

	println(&x == &x, &x == &y, &x == nil)
}