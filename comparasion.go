package main

import "fmt"

func main()  {
	var name1 = "boruto"
	var name2 = "boruto"
	var result bool = name1 == name2
	fmt.Println("output result :", result)

	var value1  = 10
	var value2 = 100
	fmt.Println(value1 > value2) // lebih dari
	fmt.Println(value1 < value2) // kurang dari
	fmt.Println(value1 == value2) // sama dengan
	fmt.Println(value1 != value2) // tidak sama
}