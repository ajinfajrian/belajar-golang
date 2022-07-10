package main

import "fmt"

func main()  {
	type noKTP string
	type married bool

	var marriedStatus married = true
	var ktpNO noKTP = "21092091"
	fmt.Println(ktpNO)
	fmt.Println(marriedStatus)
}