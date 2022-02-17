package main

func main() {
	var p = f()
	println(p)

	//each call return disctinct value
	println(f() == f())
}

func f() *int {
	v := 1
	return &v
}
