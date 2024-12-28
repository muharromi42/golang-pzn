package main

import "fmt"

func main() {
	name := "romaya"

	switch name {
	case "romay":
		fmt.Println("hello romay")
	case "albert":
		fmt.Println("hello albert")
	default:
		fmt.Println("halo boleh kenalan")
	}

	fmt.Println("=====================")

	// switch short statement
	switch lenght := len(name); lenght > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	fmt.Println("=====================")

	// switch tanpa kondisi
	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("nama terlalu panjang")
	case length2 > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
