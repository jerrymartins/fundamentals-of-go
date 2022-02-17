package main

var global *int

// x escaping from f
func f1() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
