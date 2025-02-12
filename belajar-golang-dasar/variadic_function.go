package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 20))

	//	jika sudah ada slice nya maka untuk mengkonversi menjadi varargs tambahkan titik 3
	numbers := []int{12, 4234, 313}
	fmt.Println(sumAll(numbers...))
}
