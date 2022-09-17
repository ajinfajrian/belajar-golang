package main

import "fmt"

func main()  {
	var nilaiAkhir = 80
	var nilaiAbsensi = 90
	//var totalNilai bool = nilaiAkhir >= 80
	//var totalAbsen bool = nilaiAbsensi >= 80

	//var total = totalAbsen && totalNilai
	//fmt.Println(total)
	fmt.Println(nilaiAkhir >= 80 && nilaiAbsensi >= 80)
}