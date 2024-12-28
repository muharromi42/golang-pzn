package main

import "fmt"

func main() {
	name := "albert"

	if name == "romay" {
		fmt.Println("hello",name)
	} else if name == "albert" {
		fmt.Println("hello albert")
	} else {
		fmt.Println("halo",name,"boleh kenalan?")
	}

	fmt.Println("===================================")

	// if short statement
	nama := "putra ardhana"

	// di if ini membuat statement dulu baru kondisinya
	if length := len(nama); length > 10 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}