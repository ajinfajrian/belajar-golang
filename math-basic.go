package main

import (
	"fmt"
)

func main()  {
	var result int
	result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println("c result: ",c)

	// augmented assignment
	c += 10 // setara dengan c = c + 10
	fmt.Println("augmented assignment: ", c)

	// unary operator
	c++ // setara dengan c = c + 1
	fmt.Println("i result: ", c)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}