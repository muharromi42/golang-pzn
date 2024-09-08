package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsen = 80


	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsen bool = nilaiAbsen > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsen

	fmt.Println(lulus);
}