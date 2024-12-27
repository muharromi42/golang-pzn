package main

import "fmt"

func main() {
	// ini cara lama menginisialisasikan dulu baru di isi
	// var person map[string]string = map[string]string{}
	// person["name"] = "romay"
	// person["address"] = "palembang"

	// ini cara baru langsung di isi nilainya
	person := map[string]string{
		"name" : "romay",
		"address" : "palembang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println("===================================================")
	
	book := make(map[string]string)
	book["title"] = "100 malam"
	book["author"] = "romay"
	book["wrong"] = "oopss"

	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)
}