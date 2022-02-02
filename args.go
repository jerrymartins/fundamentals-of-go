package main

import "os"

func main() {
	var s, step string

	for i := 1; i < len(os.Args); i++ {
		s += step + os.Args[i]
		step = " "
	}

	println(s)
}
