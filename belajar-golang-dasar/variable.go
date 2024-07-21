package main

import "fmt"

func main() {
	var nama string
	var angka int64
	var bool = true
	var nama2 = "putra"
	nama3 := "ardhana3"

	nama = "romi"

	fmt.Println(nama)

	nama = "ardhana"

	fmt.Println(nama)

	angka = 20

	fmt.Println(angka)
	fmt.Println(bool)
	fmt.Println(nama2)
	fmt.Println(nama3)


	// multiple variable

	var (
		firstname = "romi"
		lastname = "arhdana"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}