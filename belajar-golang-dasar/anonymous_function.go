package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	//	membuat function di dalam variabel tanpa nama
	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("anjing", blacklist)

	//	bisa disederhanakan seperti ini
	registerUser("anjing", func(name string) bool {
		return name == "romay"
	})
}
