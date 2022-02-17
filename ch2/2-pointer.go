package main

func main() {
	x := 1
	p := &x

	println(*p)

	*p = 2

	println(x)
}
