package main

import "fmt"

func main()  {
	const firstname string = "Boruto"
	const lastname = "Uzumaki"
	const value = 1000

	fmt.Println(firstname, lastname)
	fmt.Println(value)

	const( // multiple constant
		third string = "Uchiha sasuke"
		fourth = "Kiba inazuka"
	)
	fmt.Println(third, fourth)
}