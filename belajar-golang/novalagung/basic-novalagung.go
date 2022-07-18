package main

import (
	"fmt"
)

func main() {
	one := "satu"
	Str, Intgr, float, bole := "string", 1, 2.2, true // interference multi tipe data
	fmt.Println(one)
	fmt.Println(Str, Intgr, float, bole)
	var (
		left  = false
		right = true
	)
	var leftAndRight = left && right
	fmt.Printf("left && right\t(%t)\n", leftAndRight)
}
