package main

import (
	"fmt"
)

func main() {
	var firstName string
	var lastName string
	fmt.Print("First name: ")
	fmt.Scan(&firstName)
	fmt.Print("Middle name: ")
	fmt.Scan(&lastName)
	n := firstName + " " + lastName
	fmt.Println("Your Full name is", n)
}
