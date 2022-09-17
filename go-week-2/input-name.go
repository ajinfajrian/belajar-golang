package main

import (
	f "fmt"
)

func main() {
	f.Println("Nama depan: ")
	var first string

	f.Scanln(&first)
	f.Println("Nama kedua: ")
	var second string
	f.Scanln(&second)
	f.Print("Nama Panjangmu adalah ")

	f.Print(first + " " + second + "\n")
}
