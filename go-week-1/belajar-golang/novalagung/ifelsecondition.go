package main

import (
	"fmt"
)

func main()  {
	var nilaiSiswa = 4

	if nilaiSiswa == 10 {
		fmt.Println("Lulus dengan Sempurna")
	} else if nilaiSiswa > 5{
		fmt.Println("Lulus")
	} else if nilaiSiswa == 4 {
		fmt.Println("Hampir Lulus")
	} else {
		fmt.Println("Tidak lulus")
	}
}