package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "muharromi"
	middlename = "putra"
	lastname = "ardhana"

	return firstname, middlename, lastname
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
