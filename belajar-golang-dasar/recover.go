package main

import "fmt"

func endApp2() {
	fmt.Println("end App")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp2(error bool) {
	defer endApp2()
	if error {
		panic("upss error")
	}
}

func main() {
	runApp2(true)
	fmt.Println("muharromi")
}
