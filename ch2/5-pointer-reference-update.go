package main

func main() {
	v := 1
	incr(&v)

	println(v)
	println(incr(&v))
}

func incr(p *int) int {
	*p++
	return *p
}
