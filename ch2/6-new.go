package main

func main() {
	p := new(int)
	println(*p)
	*p = 2
	println(*p)
}
