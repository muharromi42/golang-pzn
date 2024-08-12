package main

import "fmt"

func main() {
	var a = 1
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d * e
	fmt.Println(c)


	// augmented assignment
	// fungsi untuk menambahkan ke variabel sendiri

	var i = 10
	i+= 20

	fmt.Println(i)


	// unary operator

	var j = 1

	j++
	j++
	j--

	fmt.Println(j)
}