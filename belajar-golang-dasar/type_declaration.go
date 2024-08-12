package main

import "fmt"

func main() {
	type NoKTP string

	var ktpromi NoKTP = "1212"


	// konversi tipe data
	var contoh string = "2222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpromi)
	fmt.Println(contohKtp)
}