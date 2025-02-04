package main

import "fmt"

func getHello(name string) string {
	hello := "Hello, " + name
	return hello
}

func main() {
	result := getHello("Romay")
	fmt.Println(result)

	fmt.Println(getHello("Albert"))
}
