package main

import "fmt"

func getFullName() (string, string) {
	return "romay", "ardhana"
}

func main() {
	//firstname, lastname := getFullName()
	//fmt.Println(firstname, lastname)

	//untuk mengambil data 1 saja
	firstname, _ := getFullName()
	fmt.Println(firstname)
}
