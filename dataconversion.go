package main

import "fmt"

func main()  {
	var nilai32 int32 = 32477
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // output berberda karna batas limit int8 adalah 127 dan mengulang lagi dari minimum

	var name = "Sasuke"
	var e byte = name[0]
	var eString = string(e) // deklarasi e menjadi string jadi output tidak lagi byte
	fmt.Println(name)
	fmt.Println(e) // output byte
	fmt.Println(eString) // output string yang mana name 0 = S
}