package main

import (
	"fmt"
	"log"
)

func main() {
	var nilai int
	fmt.Print("Input Nilai Anda: ")
	n, err := fmt.Scanln(&nilai)
	if err != nil {
		log.Fatal(err)
	}

	if nilai == 10 {
		fmt.Println("Selamat anda lulus")
	} else if nilai >= 5 {
		fmt.Println("Anda tetap lulus")
	} else if nilai == 4 {
		fmt.Println("Anda Hampir lulus")
	} else {
		fmt.Printf("Anda tidak lulus, nilai anda: %v \n", n)
	}
}
