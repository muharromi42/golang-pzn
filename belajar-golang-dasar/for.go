package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("ini perulangan ke", counter)
		counter++
	}

	fmt.Println("selesai")

	for hitung := 1; hitung <= 10; hitung++ {
		fmt.Println("ini perulangan ke", hitung)
	}

	names := []string{
		"romay",
		"albert",
		"kevin",
	}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println("name:", name)
	}
}
